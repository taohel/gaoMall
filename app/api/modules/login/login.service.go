package login

import (
	"context"
	"errors"
	"gaoMall/app"
	"gaoMall/app/api/modules/user"
	"gaoMall/app/models"
	"gaoMall/app/utils/token"
	"math/rand"
	"strconv"
	"time"
)

func GenPhoneCode(phone string) error {
	loginKey := "login_" + phone
	countKey := "login_count_" + phone
	ctx := context.Background()

	// 查看手机号是否频繁请求
	count := app.RedisR().Get(ctx, countKey).Val()
	if count == "" {
		count = "0"
	}
	countInt, err := strconv.Atoi(count)
	if err != nil {
		return err
	}
	if countInt > 2 {
		return errors.New("验证码请求频繁")
	}

	// 生成验证码
	var code string
	if app.Config.SMS.Enable {
		for i := 0; i < 6; i++ {
			code += strconv.Itoa(rand.Intn(10))
		}
	} else {
		code = app.Config.App.Code
	}

	// 测试账户 固定验证码
	if phone == "15012341234" {
		code = app.Config.App.Code
	}
	if phone == "15023452345" {
		code = app.Config.App.Code
	}

	// 存储验证码 以及请求次数
	pipe := app.RedisW().Pipeline()
	pipe.Incr(ctx, countKey)
	pipe.Expire(ctx, countKey, 2*time.Minute)
	pipe.Set(ctx, loginKey, code, 5*time.Minute)
	_, err = pipe.Exec(ctx)
	return err
}

func PhoneCode(phone, code string) (string, error) {
	var userID uint
	loginKey := "login_" + phone
	countKey := "login_count_" + phone

	// 核对手机号与验证码 删除redis信息
	ctx := context.Background()
	cacheCode := app.RedisR().Get(ctx, loginKey).Val()
	if cacheCode != code {
		return "", errors.New("验证码错误")
	}
	app.RedisW().Del(ctx, loginKey)
	app.RedisW().Del(ctx, countKey)

	// 查询用户是登录还是注册
	userRel, err := user.FindOneByPhone(phone)
	if err != nil {
		return "", err
	}

	if userRel.ID == 0 {
		// 用户不存在 注册
		var newUser = models.User{Phone: phone}
		err = user.Create(&newUser)
		if err != nil {
			return "", err
		}
		userID = newUser.ID
	} else {
		userID = userRel.ID
	}

	// 创建token
	return token.CreateToken(userID)
}
