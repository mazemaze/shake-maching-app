package router

import (
	"net/http"
	"shakashaka/model"

	"github.com/gin-gonic/gin"
)

func createChatHandler(ctx *gin.Context) {
	var content model.ChatContent
	_ = ctx.Param("id")
	err := ctx.BindJSON(content)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"result": true})
}

func getChatByRoomID(ctx *gin.Context) {
	_ = ctx.Param("id")
}
