package auth

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 白名单
var whitelist = []string{
	"/",
	"/login",
	"/websockets",
	"/ws",
}

// Auth 中间件
func Auth(ctx *gin.Context) {
	for _, v := range whitelist {
		if v == ctx.Request.URL.Path {
			ctx.Next()
			return
		}
	}
	session := sessions.Default(ctx)
	sessionId := ctx.Request.Header.Get("X-Session-ID")
	if sessionId == "" {
		ctx.JSON(http.StatusOK, gin.H{"code": 401, "message": "未登录"})
		ctx.Abort()
		return
	}
	userInfo := session.Get(sessionId)
	if userInfo == nil {
		ctx.JSON(http.StatusOK, gin.H{"code": 401, "message": "未登录"})
		ctx.Abort()
		return
	}
	ctx.Set("userId", userInfo)
	ctx.Next()
}
