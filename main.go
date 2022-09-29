package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"sso_gin/api"
	"sso_gin/db"
)

var (
	HandleRegister = api.HandleRegister
)

func main() {
	db.ConnectMYSQL()
	r := gin.Default()

	r.POST("/register", HandleRegister)

	err := r.Run(":3000")
	if err != nil {
		log.Fatalf("gin运行出错：%v", err)
	}
}
