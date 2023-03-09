package router

import "github.com/gin-gonic/gin"

func Router() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")

	api.GET("/user/:id", userInfoHandler)
	api.PUT("/user", updateUserInfoHandler)
	api.POST("/user/login", loginHandler)
	api.POST("/user/registration", registrationHandler)

	api.POST("/chat_room", createChatRoomHandler)
	api.GET("/chat_room/:id", chatRoomInfoHanlder)
	api.PUT("/chat_room/:id", updateChatRoomStatus)

	api.POST("/chat_room/chats", createChatHandler)
	api.GET("/chat_room/:id/chats", getChatByRoomID)

	return r
}
