package user

import (
	"fmt"
	"net/http"
	userController "warehouse/controller/user"
	"warehouse/model"

	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler(rg *gin.RouterGroup) {
	h := Handler{}

	rg.POST("/login", h.login)
	rg.POST("/users/register", h.register)

	rg.GET("/users/:id", h.getUser)
	rg.GET("/users", h.getAllUser)
}

func (h *Handler) register(c *gin.Context) {
	userInfo := model.UserInfo{}

	if err := c.ShouldBindJSON(&userInfo); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		fmt.Println("register error")
		return
	} else {
		userController.UserRegister(userInfo)
	}

	// response 200
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
	})
}

func (h *Handler) login(c *gin.Context) {
	userInfo := model.UserInfo{}

	if err := c.ShouldBindJSON(&userInfo); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	} else {
		// user login server
		userController.UserLogin(userInfo, c.ClientIP())
	}

	// response 200
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
	})
}

func (h *Handler) getUser(c *gin.Context) {
	userInfo := model.UserInfo{}

	id := c.Param("id")
	// user login server
	userController.GetUser(&userInfo, id)

	c.JSON(http.StatusOK, userInfo)
}

func (h *Handler) getAllUser(c *gin.Context) {
	userInfos := []model.UserInfo{}

	// user login server
	userController.GetAllUser(&userInfos)

	c.JSON(http.StatusOK, userInfos)
}
