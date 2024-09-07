package model

// Message 发送的消息
type Message struct {
	ID           int    `json:"id"`
	User         *User  `json:"user"`
	Content      string `json:"content"`
	SendToUserID int    `json:"sendToUserID"`
}
