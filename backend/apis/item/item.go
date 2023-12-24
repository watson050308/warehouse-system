package item

import (
	"fmt"
	"net/http"
	"warehouse/model"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	product model.ProductInfo
}

func NewHandler(rg *gin.RouterGroup) {
	h := Handler{}

	rg.GET("/items", h.getAllItem)
	rg.GET("/items/:id", h.getItem)

	rg.POST("/items", h.addItem)

	rg.DELETE("/items/:id", h.deleteItem)

}

func (h *Handler) getAllItem(c *gin.Context) {
	products := h.product.GetAllItem()
	c.JSON(http.StatusOK, products)
}

func (h *Handler) getItem(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		fmt.Println("not set id")
		return
	}

	h.product.ID = id
	h.product.GetItem()
	c.JSON(http.StatusOK, h.product)
}

func (h *Handler) addItem(c *gin.Context) {
	if err := c.ShouldBindJSON(&h.product); err != nil {
		fmt.Println("err bind json item", err.Error())
		return
	}

	if err := h.product.AddItem(); err != nil {
		fmt.Println("err insert item", err.Error())
		return
	}

	// response 200
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
	})
}

func (h *Handler) deleteItem(c *gin.Context) {
	id := c.Param("id")

	if err := h.product.DeleteItem(id); err != nil {
		return
	}

	// response 200
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
	})
}
