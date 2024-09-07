package logic

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"yy-im/domain/common"
	"yy-im/socket"
)

// WsHandle 建立websocket连接
func WsHandle(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Query("userId"))
	conn, err := socket.Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		ctx.JSON(500, gin.H{"msg": "连接失败"})
		return
	}
	defer conn.Close()
	err = socket.Connect(conn, userId)
	if err != nil {
		ctx.JSON(500, gin.H{"msg": "连接失败"})
		return
	}

	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			ctx.JSON(500, gin.H{"msg": "连接失败"})
			return
		}
		socketMsg := common.SocketMsg{}
		err = json.Unmarshal(msg, &socketMsg)
		if err != nil {
			fmt.Printf("json.Unmarshal error: %s\n", err)
		}
		fmt.Println(socketMsg)
		switch socketMsg.Action {
		case "sendMsg":
			// 发送消息
			err, _ := socket.SendMsg(ctx, socketMsg)
			if err != nil {
				fmt.Printf("sendMsg error: %s\n", err)
			}

		default:
			fmt.Printf("unknown action: %s\n", socketMsg.Action)
		}
		fmt.Printf("消息类型 %d", messageType)
	}

}
