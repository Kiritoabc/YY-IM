package logic

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"yy-im/domain"
	"yy-im/domain/common"
	"yy-im/domain/model"
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
	u, ok := check(user.UserName, user.Password)
	if !ok {
		ctx.JSON(401, gin.H{"code": 1, "msg": "用户名或密码错误"})
	}
	// session
	session := sessions.Default(ctx)
	sessionId := uuid.New().String()
	// 先删除，再保存，防止多次登录，目前不设置登出
	session.Delete(u.ID)
	session.Set(sessionId, u.ID)
	err := session.Save()
	if err != nil {
		ctx.JSON(500, gin.H{"code": 1, "msg": "session 保存失败"})
		return
	}

	ctx.JSON(200, common.LoginResp{
		Code:      0,
		SessionId: sessionId,
		UserID:    u.ID,
		UserName:  u.Name,
	})
	return
}

// check 检查用户
func check(username, password string) (*model.User, bool) {
	user := domain.GetUserByName(username)
	if user == nil {
		return nil, false
	}
	return user, user.Password == password
}
