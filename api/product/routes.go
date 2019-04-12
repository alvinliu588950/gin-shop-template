package product

import (
	"github.com/gin-gonic/gin"
	"teashop/api/product/handler"
)


func Routes(r *gin.Engine) {
	r.GET("/product/list", handler.List())
	r.POST("/product/add", handler.Add())
	r.PATCH("/product/update", handler.Update())
	r.DELETE("/product/delete", handler.Delete())
}