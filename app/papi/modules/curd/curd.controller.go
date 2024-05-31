package curd

import (
	"gaoMall/app/models"
	_ "gaoMall/app/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

var modelMap = map[string]interface{}{
	"user":    &models.User{},
	"article": &models.Article{},
}

func FindCo(ctx *gin.Context) {
	model := ctx.Param("model")
	if modelMap[model] == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": "false",
			"message": "模块不存在",
		})
		return
	}

	data := Find(modelMap[model])
	ctx.JSON(http.StatusOK, gin.H{
		"success": "true",
		"data":    data,
	})
}
