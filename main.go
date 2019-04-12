package main


import (
	"github.com/gin-gonic/gin"
	"teashop/api/product"
)

func main() {
	r := gin.Default()

	product.Routes(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}