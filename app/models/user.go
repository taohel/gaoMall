package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Phone        string `json:"phone" gorm:"unique"`
	Username     string `json:"username" gorm:"unique"`
	Email        string `json:"email" gorm:"unique"`
	WechatOpenId string `json:"wechatOpenId" gorm:"unique"`
	Password     string `json:"password"`
	Nickname     string `json:"nickname" gorm:"default:新用户"`
	Avatar       string `json:"avatar" `
}
