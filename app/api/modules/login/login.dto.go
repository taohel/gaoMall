package login

type GetPhoneCodeReq struct {
	Phone string `form:"phone" binding:"required,min=11,max=11" example:"15012341234"`
}

type PhoneCodeReq struct {
	Phone string `json:"phone" binding:"required,min=11,max=11" example:"15012341234"`
	Code  string `json:"code" binding:"required,min=6,max=6" example:"123456"`
}
