package routes

import (
	"introduccion_go/src/products/infrastucture/dependencies"

	"github.com/gin-gonic/gin"
)

func Routes (r *gin.Engine){

	productsRoutes := r.Group("/products")
	

	createProduct := dependencies.CreateProductController()
	getAllProduct := dependencies.GetAllProductsController()
	updateProduct := dependencies.UpdateProductsController()
	deleteProduct := dependencies.DeleteProductsController()
	

	productsRoutes.POST("/add", createProduct.Run)
	productsRoutes.GET("/", getAllProduct.Run)
	productsRoutes.PUT("/update/:id", updateProduct.Run)
	productsRoutes.DELETE("/delete/id", deleteProduct.Run)

}