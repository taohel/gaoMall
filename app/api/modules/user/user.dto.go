package user

type LoginReq struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type InfoRes struct {
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}
