package main

import (
	"log"
	"sso_gin/api"
	"sso_gin/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.ConnectMYSQL()
	r := gin.Default()

	api.SetupRouter(r)

	err := r.Run(":3000")
	if err != nil {
		log.Fatalf("gin运行出错：%v", err)
	}
}
