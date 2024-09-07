package socket

import (
	"context"
	"errors"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
	"yy-im/domain/common"
)

// WebSocketConn 记录所有的websocket连接
var WebSocketConn = make(map[int]*websocket.Conn)
var mutex = sync.Mutex{}

// Upgrader 升级HTTP连接为websocket连接
var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许所有来源的请求
	},
}

// Connect 建立连接
func Connect(conn *websocket.Conn, userId int) error {
	mutex.Lock()
	WebSocketConn[userId] = conn
	mutex.Unlock()
	return nil
}

// SendMsg 发送消息
func SendMsg(ctx context.Context, msg common.SocketMsg) (error, bool) {
	conn := getConn(msg.TargetUserId)
	if conn == nil {
		// todo: 暂时其他用户未登录时，不存储消息
		return errors.New("其他用户未登录"), false
	}
	if err := conn.WriteMessage(websocket.TextMessage, []byte(msg.Message)); err != nil {
		return err, false
	}
	return nil, true
}

// getConn 获取连接
func getConn(id int) *websocket.Conn {
	return WebSocketConn[id]
}
