package router

import (
	"log"
	"net/http"
	"shakashaka/model"

	"github.com/gin-gonic/gin"
)

func createChatHandler(ctx *gin.Context) {
	var content model.ChatContent
	err := ctx.BindJSON(content)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	db := DBInit()
	err = db.Create(&content).Error
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"result": false})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": true})
}

func getChatByRoomID(ctx *gin.Context) {
	id := ctx.Param("id")
	db := DBInit()
	var chats []model.ChatContent
	result := db.Where("room_id = ?", id).Find(&chats)
	if result.Error != nil {
		log.Println(result.Error)
		ctx.JSON(http.StatusBadRequest, gin.H{"result": false})
		return
	}
	ctx.JSON(http.StatusOK, chats)

}
