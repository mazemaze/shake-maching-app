package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique" json"username"`
	Password string `json"password"`
}

type ChatRoom struct {
	gorm.Model
	User1 string `json:"user_1"`
	User2 string `json:"user_2"`
}

type ChatContent struct {
	gorm.Model
	RoomId  string `json:"room_id"`
	UserId  string `json:"user_id"`
	Content string `json:"content"`
}

type Score struct {
	gorm.Model
	RoomId string `json:"room_id"`
	UserId string `json:"user_id"`
	Score  int64  `json:"score"`
}
