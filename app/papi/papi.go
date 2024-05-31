package papi

import (
	"gaoMall/app"
	"gaoMall/app/api/middleware"
	"gaoMall/app/papi/routes"
	"github.com/gin-gonic/gin"
)

func NewServer() {
	r := gin.Default()
	r.Use(middleware.Cors())
	routes.RegisterRoute(r)
	err := r.Run(app.Config.PanelApp.Addr)
	if err != nil {
		panic(err)
	}
}
