package routes

import (
	"introduccion_go/src/products/infrastucture/controllers"

	"github.com/gin-gonic/gin"
)

func routes (r *gin.Engine){
	productsRoutes := r.Group("/products")
	{
		productsRoutes.POST("/add", controllers.CreateProduct)
	}
}