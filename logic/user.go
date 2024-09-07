package logic

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"yy-im/domain"
)

// GetAllUser 获取所有用户
func GetAllUser(ctx *gin.Context) {
	userList := domain.GetAllUser()
	userId, _ := strconv.Atoi(ctx.Query("userId"))
	for i := 0; i < len(userList); i++ {
		if userList[i].ID == userId {
			userList = append(userList[:i], userList[i+1:]...)
		}
	}
	ctx.JSON(200, userList)
}
