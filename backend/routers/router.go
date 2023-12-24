package routers

import (
	"warehouse/apis/item"
	"warehouse/apis/user"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery())
	InitAdmin(router)
	return router
}

func InitAdmin(router *gin.Engine) {
	rg := router.Group("/api")

	user.NewHandler(rg)
	item.NewHandler(rg)
}
