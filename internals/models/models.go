package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Username  string    `json:"username" gorm:"unique"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	Sessions  []Session `json:"sessions" gorm:"foreignKey:UserIDs"`
}

type Session struct {
	ID       uint      `json:"id" gorm:"primary_key"`
	SocketId string    `json:"socket_id"`
	UserIDs  []uint    `json:"user_ids" gorm:"type:integer[]; optional"`
	Messages []Message `json:"messages" gorm:"foreignKey:SessionID"`
}

type Message struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Text      string `json:"text"`
	SessionID uint   `json:"session_id"`
	// Session   Session   `json:"session" gorm:" foreignKey:SessionID;references:ID"`
	UserID    uint      `json:"user_id" gorm:"optional"`
	UserType  string    `json:"user_type" gorm:"optional"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}
