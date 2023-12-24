package main

import (
	"warehouse/db"
	"warehouse/routers"
)

func main() {

	var mysql *db.MysqlDB
	mysql.Init()

	r := routers.Init()
	r.Run(":8080")
}
