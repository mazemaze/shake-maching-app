package router

import "github.com/gin-gonic/gin"

func Router() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")

	api.GET("/user/:id", userInfoHandler)
	api.POST("/user/login", loginHandler)
	api.POST("/user/registration", registrationHandler)

	api.POST("/chat_room", createChatRoomHandler)
	api.GET("/chat_room/:id", chatRoomInfoHanlder)

	api.POST("/chat_room/:id", createChatHandler)
	api.GET("/chat_room/:id/chats", getChatByRoomID)
	return r
}
