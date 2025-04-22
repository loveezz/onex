package main

import (
	"one/controller"
	"one/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	r := gin.Default()

	


	testrouter := r.Group("/api/products")

	{
		testrouter.GET("/", controller.InitProducts)
		testrouter.GET("/:id", controller.GetProduct)
		testrouter.POST("/", controller.CreateProduct)
		testrouter.PUT("/:id", controller.UpdateProduct)
		testrouter.DELETE("/:id", controller.DeleteProduct)
	}

	r.Run(":8080")
		
	
}