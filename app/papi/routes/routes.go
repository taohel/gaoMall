package routes

import (
	"gaoMall/app"
	"gaoMall/app/papi/modules/curd"
	"github.com/gin-gonic/gin"
)

func RegisterRoute(r *gin.Engine) {
	rg := r.Group("/api")
	rg.GET("/hi", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": app.Config.PanelApp.Name + " app running..."})
	})

	curdRoutes := r.Group("/api/curd")
	curdRoutes.GET("/:model", curd.FindCo)
}
