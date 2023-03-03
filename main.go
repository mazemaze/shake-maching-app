package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"shakashaka/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wshandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to set websocket upgrade: %+v", err)
		return
	}
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	//何か受け取ってそのまま返すパターン

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		conn.WriteMessage(t, msg)
	}
}

func main() {
	fmt.Println("Hi")

	r := gin.Default()

	api := r.Group("/api")

	authrized := api.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo": "test",
	}))

	authrized.GET("/secrets", func(ctx *gin.Context) {

		ctx.JSON(http.StatusOK, "Hi!")

	})

	api.GET("/registration", func(r *gin.Context) {

	})

	api.GET("/ws", func(c *gin.Context) {
		wshandler(c.Writer, c.Request)
	})
	go clientStart()

	api.POST("/chatroom", func(c *gin.Context) {
		var chatRoom *model.ChatRoom
		err := c.BindJSON(&chatRoom)
		if err != nil {
			log.Fatalln("Failed to parse the json: " + err.Error())
			return
		}
	})

	r.Run(":8080")
}

func clientStart() {
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/api/ws"}
	log.Printf("connecting to %s", u.String())
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Println("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go receive(c, done)
	waitloop(c, done)
}

func receive(c *websocket.Conn, done chan struct{}) {
	defer close(done)
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		log.Printf("recv: %s", message)
	}
}

func waitloop(c *websocket.Conn, done chan struct{}) {
	//停止用
	//interrupt := make(chan os.Signal, 1)
	//signal.Notify(interrupt, os.Interrupt)
	for {
		select {
		case <-done:
			log.Println("done")
			return
			/*
				case <-interrupt:
					log.Println("interrupt")
					err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
					if err != nil {
						log.Println("write close:", err)
						return
					}
					select {
					case <-done:
					case <-time.After(time.Second):
					}
					return
			*/
		}
	}
}
