package handler

import (
	"github.com/gin-gonic/gin"
	"teashop/api/product/model"
)


func List() gin.HandlerFunc {
	return func (c *gin.Context) {
		p := model.Product{}

		result, _ := p.GetAll()

		c.JSON(200, gin.H{
			"products": result,
		})
	}
}