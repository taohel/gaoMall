package user

import (
	"errors"
	"gaoMall/app"
	"gaoMall/app/models"
	"gorm.io/gorm"
)

func Create(user *models.User) (err error) {
	return app.DBW().Create(&user).Error
}

func FindOne(user *models.User) (models.User, error) {
	var u models.User
	err := app.DBR().Where(&user).First(&u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	return u, err
}

func FindOneByPhone(phone string) (models.User, error) {
	var u models.User
	err := app.DBR().Where(&models.User{Phone: phone}).First(&u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	return u, err
}
