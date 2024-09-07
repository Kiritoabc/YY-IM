package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"log"
	"yy-im/logic"
)
import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	store := cookie.NewStore([]byte("secret-key"))

	router.Use(gin.Logger()).Use(sessions.Sessions("mysession", store))

	router.LoadHTMLGlob("templates/*")
	// 发送到具体用户
	router.GET("/send", logic.SendHandle)
	// 发送到所有人
	router.GET("/echo", logic.EchoHandle)

	router.GET("/login", logic.Login)
	log.Fatal(router.Run(":80"))
}
