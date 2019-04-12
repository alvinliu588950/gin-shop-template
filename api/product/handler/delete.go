package handler

import (
	"teashop/api/product/model"
	"github.com/gin-gonic/gin"
)

type DeleteRequest struct {
	Id uint `form:"id" json:"id" xml:"id"  binding:"exists"`
}

func Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body DeleteRequest
		
		if err := c.ShouldBind(&body); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		p := model.Product{
			Id: body.Id,
		}

		err := p.Delete()

		if err == nil {
			c.JSON(200, gin.H{
				"message": "delete success",
			})
		}

	}
}
