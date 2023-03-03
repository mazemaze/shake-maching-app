package model

import "time"

type User struct {
	Id        string    `json:"id"`
	Username  string    `json"username"`
	Password  string    `json"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ChatRoom struct {
	Id        string `json:"id"`
	User1     string `json:"user_1"`
	User2     string `json:"user_2"`
	CreatedAt string `json:"created_at"`
}

type ChatContent struct {
	Id        string `json:"id"`
	RoomId    string `json:"room_id"`
	UserId    string `json:"user_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

type Score struct {
	Id     string `json:"id"`
	RoomId string `json:"room_id"`
	UserId string `json:"user_id"`
	Score  int64  `json:"score"`
}
