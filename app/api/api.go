package api

import (
	"gaoMall/app"
	"gaoMall/app/api/middleware"
	"gaoMall/app/api/routes"
	"gaoMall/app/models"
	"github.com/gin-gonic/gin"
)

func NewServer() {
	err := app.DBW().AutoMigrate(&models.User{})
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.Use(middleware.Cors())
	r.Static("/public", "static/public")
	r.Static("/uploads", "static/uploads")
	routes.RegisterRoute(r)
	err = r.Run(app.Config.App.Addr)
	if err != nil {
		panic(err)
	}
}
