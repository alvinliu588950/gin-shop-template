package handler

import (
	"teashop/api/product/model"
	"github.com/gin-gonic/gin"
)

type UpdateRequest struct {
	Id uint `form:"id" json:"id" xml:"id"  binding:"exists"`
	Name string `form:"name" json:"name" xml:"name"  binding:"exists"`
	Price float32 `form:"price" json:"price" xml:"price"  binding:"exists"`
}

func Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body UpdateRequest
		
		if err := c.ShouldBind(&body); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		p := model.Product{
			Id: body.Id,
			Name: body.Name,
			Price: body.Price,
		}

		err := p.Update()

		if err == nil {
			c.JSON(200, gin.H{
				"message": "update success",
			})
		}

	}
}
