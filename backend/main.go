package main

import (
	"warehouse/apis/user"
	"warehouse/db"

	"github.com/gin-gonic/gin"
)

func main() {

	var mysql *db.MysqlDB
	mysql.Init()

	r := gin.New()
	r.Use(gin.Recovery())

	rg := r.Group("/api")

	user.NewHandler(rg)
	r.Run(":8080")
}
