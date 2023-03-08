package router

import "github.com/gin-gonic/gin"

func createChatRoomHandler(ctx *gin.Context) {
}

func chatRoomInfoHanlder(ctx *gin.Context) {
	_ = ctx.Param("id")
}
