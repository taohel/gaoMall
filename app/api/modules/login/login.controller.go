package login

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPhoneCodeCo(ctx *gin.Context) {
	var req GetPhoneCodeReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := GenPhoneCode(req.Phone)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"success": false, "message": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"success": true})
	}
}

func PhoneCodeCo(ctx *gin.Context) {
	var req PhoneCodeReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := PhoneCode(req.Phone, req.Code)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"success": false, "message": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"success": true, "data": token})
	}
}
