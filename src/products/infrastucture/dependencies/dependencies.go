package dependencies

import (
	"fmt"
	"introduccion_go/src/Core"
	"introduccion_go/src/products/application"
	"introduccion_go/src/products/infrastucture"
	"introduccion_go/src/products/infrastucture/controllers"
)

var (mySQL infraestructure.MySQL)

func Init () {
	conn := core.GetDBPool() 
	if conn.Err != "" {
		fmt.Println("Error de servidor", conn.Err)
		return
	}

	mySQL = *infraestructure.NewMySQL(conn.DB)
}

//Casos de uso quye aplicamos acá
func CreateProductController() *controllers.CreateProductController {
	ucCreateProduct := application.NewCreateProduct(&mySQL)
	return controllers.NewCreateProductController(ucCreateProduct)
}

func GetAllProductsController() *controllers.GetAllProductController{
	ucGetAllProducts := application.NewGetAllProduct(&mySQL)
	return controllers.NewGetAllProductController(ucGetAllProducts)
}

func UpdateProductsController() *controllers.UpdateProductByIdController{
	ucUpdateProductsController := application.NewUpdateProduct(&mySQL)
	return controllers.NewUpdateProductByIdController(ucUpdateProductsController)
}

func DeleteProductsController() *controllers.DeleteProductByIdController{
	ucDeleteProductsController := application.NewDeleteProduct(&mySQL)
	return controllers.NewDeleteProductByIdController(ucDeleteProductsController)
}


