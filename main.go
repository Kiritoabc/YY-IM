package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"log"
	"yy-im/pkg/auth"

	"yy-im/logic"
)

func main() {
	router := gin.Default()
	store := cookie.NewStore([]byte("secret-key"))

	router.Use(gin.Logger()).Use(sessions.Sessions("mysession", store))

	router.LoadHTMLGlob("templates/*")

	router.Use(auth.Auth)

	router.GET("/ws", logic.WsHandle)

	router.POST("/login", logic.Login)
	router.GET("/user/all", logic.GetAllUser)

	router.GET("/websockets", func(ctx *gin.Context) {
		ctx.HTML(200, "websockets.html", nil)
	})
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "login.html", nil)
	})
	log.Fatal(router.Run(":80"))
}
