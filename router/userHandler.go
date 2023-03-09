package router

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"shakashaka/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func userInfoHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	var user model.User
	db := DBInit()
	err := db.First(&user, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func updateUserInfoHandler(ctx *gin.Context) {
	var user model.User
	db := DBInit()
	err := db.Save(&user).Error
	if err != nil {
		log.Println("Failed to process userinfo update")
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	err = db.Where("id = ?", user.ID).First(&user).Error
	if err != nil {
		log.Println("Failed to process userInfo get")
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func registrationHandler(ctx *gin.Context) {
	var user model.User
	err := ctx.BindJSON(&user)
	if err != nil {
		log.Println("Failed to bind user request")
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"result": false})
		return
	}

	db := DBInit()
	result := db.Create(&user)
	if result.Error != nil {
		log.Println(result.Error)
		ctx.JSON(http.StatusBadRequest, gin.H{"result": false})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"result": true})
}

func loginHandler(ctx *gin.Context) {
	var request model.User
	err := ctx.BindJSON(request)
	if err != nil {
		log.Println("Failed to process user login")
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	var user model.User
	db := DBInit()
	err = db.Where("username = ?", request.Username).Where("password = ?", request.Password).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	} else if errors.Is(err, gorm.ErrRegistered) {
		log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, "テスト")
		return
	}
	fmt.Println(err)
	ctx.JSON(http.StatusOK, user)
}
