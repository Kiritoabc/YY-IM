package domain

// Message 发送的消息
type Message struct {
	ID           int    `json:"id"`
	User         *User  `json:"user"`
	Content      string `json:"content"`
	SendToUserID int    `json:"sendToUserID"`
}

// MsgDB 存储消息
var MsgDB []Message = []Message{}
