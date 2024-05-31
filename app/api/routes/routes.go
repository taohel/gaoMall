package routes

import (
	"gaoMall/app"
	"gaoMall/app/api/middleware"
	"gaoMall/app/api/modules/login"
	"gaoMall/app/api/modules/user"
	"github.com/gin-gonic/gin"
)

func RegisterRoute(r *gin.Engine) {
	rg := r.Group("/api")
	rg.GET("/hi", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": app.Config.App.Name + " app running..."})
	})

	rg.GET("/user/login/phone-code", login.GetPhoneCodeCo)
	rg.POST("/user/login/phone-code", login.PhoneCodeCo)

	rg.POST("/user/login", user.LoginCo)
	rg.POST("/user/register", user.RegisterCo)

	rg.Use(middleware.Auth())
	rg.GET("/user/info", user.GetInfoCo)
}
