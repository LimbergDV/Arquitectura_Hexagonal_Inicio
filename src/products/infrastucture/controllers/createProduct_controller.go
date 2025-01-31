package controllers

import (
	"fmt"
	"introduccion_go/src/products/application"
	"introduccion_go/src/products/domain"
	infraestructure "introduccion_go/src/products/infrastucture"
	"introduccion_go/src/products/infrastucture/routes/validators"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateProductController struct {
	app *application.CreateProduct
}

//Este constructor permite inyectar la dependencia del caso del uso (application.CreateProduct) al controlador.
func NewCreateProductController()  *CreateProductController {
	mysql := infraestructure.GetMySQL()
	app := application.NewCreateProduct(mysql)
	return &CreateProductController{app: app}
}

func (ctrl *CreateProductController) Run(c *gin.Context){
	var products domain.Product

	if err := c.ShouldBindJSON(&products); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status":false, "error": "Datos ivalidos" + err.Error()})
		return 
	}

	if err := validators.CheckProduct(products); err != nil {
		c.JSON(http.StatusBadRequest, gin.H {"status": false, "error": "Datos ivalidos" + err.E})
	} 

	rowsAffected, err := ctrl.app.Run(products)

	if err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	return

	if rowsAffected == 0 {
		fmt.Print("hola")
	}
	} else {
		c.JSON(http.StatusCreated, gin.H {"mensaje": "Producto creado"})
		c.JSON(http.StatusOK, products)
	}
}


