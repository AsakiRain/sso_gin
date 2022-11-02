package main

import (
	"log"
	"sso_gin/api"
	"sso_gin/db"
	"sso_gin/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	db.SetupMYSQL()
	db.SetupGoCache()
	r := gin.Default()
	r.Use(middleware.Cors())
	api.SetupRouter(r)

	err := r.Run(":3000")
	if err != nil {
		log.Fatalf("!!GIN运行出错：%v", err)
	}
}
