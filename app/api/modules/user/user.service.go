package user

import (
	"gaoMall/app/models"
	"gorm.io/gorm"
)

func GetInfo(userID uint) (InfoRes, error) {
	var i InfoRes
	var params = models.User{
		Model: gorm.Model{ID: userID},
	}
	userRel, err := FindOne(&params)
	if err != nil {
		return i, err
	}

	i.Nickname = userRel.Nickname
	i.Avatar = userRel.Avatar
	return i, nil
}

func Login(username, password string) (models.User, error) {
	var u models.User
	var params = models.User{
		Username: username,
		Password: password,
	}

	u, err := FindOne(&params)
	return u, err
}

func Register(username, password string) error {
	user := models.User{
		Username: username,
		Password: password,
	}
	return Create(&user)
}
