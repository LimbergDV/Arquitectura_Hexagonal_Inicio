package routes

import (
	"introduccion_go/src/products/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func Routes (r *gin.Engine){

	productsRoutes := r.Group("/products") 
	{

	productsRoutes.POST("/", controllers.NewCreateProductController().Run)
	productsRoutes.GET("/", controllers.NewGetAllProductController().Run)
	productsRoutes.PUT("/:id", controllers.NewUpdateProductByIdController().Run)
	productsRoutes.DELETE("/:id", controllers.NewDeleteProductByIdController().Run)

	}
}