package user

import (
	"gaoMall/app/api/middleware"
	"gaoMall/app/utils/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetInfoCo(ctx *gin.Context) {
	userPayload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*token.Payload)
	info, err := GetInfo(userPayload.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{"success": true, "data": info})
	return
}

func LoginCo(ctx *gin.Context) {
	var req LoginReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rel, err := Login(req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rel.ID == 0 {
		ctx.JSON(http.StatusOK, gin.H{"success": false})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true, "data": rel})
	return
}

func RegisterCo(ctx *gin.Context) {
	var req LoginReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := Register(req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true})
	return
}
