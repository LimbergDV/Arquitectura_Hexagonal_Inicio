package routes

import (
	"introduccion_go/src/products/infrastucture/controllers"

	"github.com/gin-gonic/gin"
)

func Routes (r *gin.Engine){

	productsRoutes := r.Group("/products") 
	{

	productsRoutes.POST("/add", controllers.NewCreateProductController().Run)
	productsRoutes.GET("/", controllers.NewGetAllProductController().Run)
	productsRoutes.PUT("/update/:id", controllers.NewDeleteProductByIdController().Run)
	productsRoutes.DELETE("/delete/id", controllers.NewDeleteProductByIdController().Run)

	}
}