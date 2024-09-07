package logic

import (
	"github.com/gin-gonic/gin"
	"yy-im/domain/common"
)

// Login 用户登录接口
func Login(ctx *gin.Context) {
	// 登录逻辑
	user := &common.LoginReq{}
	if err := ctx.ShouldBind(user); err != nil {
		ctx.JSON(400, gin.H{"code": 1, "msg": err.Error()})
		return
	}

	// 判断用户是否存在

	ctx.HTML(200, "websockets.html", nil)
}
