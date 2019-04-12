package handler

import (
	"github.com/gin-gonic/gin"
	"teashop/api/product/model"
)

type AddRequest struct {
	Name string `form:"name" json:"name" xml:"name"  binding:"exists"`
	Price float32 `form:"price" json:"price" xml:"price"  binding:"exists"`
}

func Add() gin.HandlerFunc {
	return func (c *gin.Context) {
		var body AddRequest
		
		if err := c.ShouldBind(&body); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		p := model.Product{
			Name: body.Name, 
			Price: body.Price,
		}

		p.Create()

		c.JSON(200, gin.H{
			"message": "create success",
		})
	}
}