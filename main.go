package main

import (
	"log"
	"net/http"
	"net/url"
	"shakashaka/model"
	"shakashaka/router"
	"time"

	"github.com/gorilla/websocket"
	"gorm.io/gorm"
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

	db := router.DBInit()

	err := migration(db)

	if err != nil {
		panic("Failed to migrate database")
	}
	router := router.Router()

	router.Run(":8080")
}

func migration(db *gorm.DB) error {
	err := db.AutoMigrate(
		&model.ChatContent{},
		&model.ChatRoom{},
		&model.Score{},
		&model.User{},
	)
	if err != nil {
		return err
	}
	return nil
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
