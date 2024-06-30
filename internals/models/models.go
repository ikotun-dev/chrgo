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
	UserIDs  []uint    `json:"user_ids" gorm:"type:integer[]"`
	Messages []Message `json:"messages" gorm:"foreignKey:ID;references:ID"`
}

type Message struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Text      string    `json:"text"`
	SessionID uint      `json:"session_id"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
