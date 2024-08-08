package models

import "time"

type User struct {
	ID                uint      `json:"id" gorm:"primary_key"`
	Email             string    `json:"email" gorm:"unique"`
	Password          string    `json:"password"`
	CreatedAt         time.Time `json:"created_at"`
	Verified          bool      `json:"verified" gorm:"default:false"`
	VerificationToken string    `json:"verification_token"`
	Sessions          []Session `json:"sessions" gorm:"foreignKey:UserIDs"`
}

type Message struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Content   string    `json:"content"`
	SessionID uint      `json:"session_id"`
	UserID    uint      `json:"user_id" gorm:"optional"`
	UserType  string    `json:"user_type" gorm:"optional"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

type Chatbot struct {
	ID                uint      `json:"id" gorm:"primary_key"`
	Name              string    `json:"name"`
	AssistantID       string    `json:"assistant_id"`
	UserID            uint      `json:"user_id" gorm:"foreignKey:ID;references:UserID"`
	UserInstruction   string    `json:"user_instruction"`
	CustomInstruction string    `json:"custom_instruction"`
	CreatedAt         time.Time `json:"created_at" gorm:"autoCreateTime"`
}

type Session struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	SocketId  string    `json:"socket_id"`
	ThreadID  string    `json:"thread_id"`
	ChatbotID uint      `json:"chatbot_id"`
	UserIDs   []uint    `json:"user_ids" gorm:"type:integer[]; optional"`
	Messages  []Message `json:"messages" gorm:"foreignKey:SessionID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

// Session   Session   `json:"session" gorm:" foreignKey:SessionID;references:ID"`
