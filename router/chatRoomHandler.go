package router

import (
	"errors"
	"log"
	"net/http"
	"shakashaka/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func createChatRoomHandler(ctx *gin.Context) {
	var chatRoom model.ChatRoom
	err := ctx.BindJSON(&chatRoom)
	if err != nil {
		log.Println("Failed to bind create chat room request: " + err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"result": false})
		return
	}
	db := DBInit()
	chatRoom.Status = 0
	result := db.Create(&chatRoom)
	if result.Error != nil {
		log.Println(result.Error)
		ctx.JSON(http.StatusBadRequest, gin.H{"result": false})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"result": true})
}

func chatRoomInfoHanlder(ctx *gin.Context) {
	id := ctx.Param("id")
	var chatRoom model.ChatRoom
	db := DBInit()
	err := db.First(&chatRoom, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, chatRoom)
}

func updateChatRoomStatus(ctx *gin.Context) {
	id := ctx.Param("id")
	db := DBInit()
	err := db.Where("id = ?", id).Update("status", 1).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, nil)
}
