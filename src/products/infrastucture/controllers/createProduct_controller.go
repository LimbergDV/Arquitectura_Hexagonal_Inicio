package controllers

import (
	"introduccion_go/src/products/application"
	"introduccion_go/src/products/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateProductController struct {
	uc *application.CreateProduct
}

//Este constructor permite inyectar la dependencia del caso del uso (application.CreateProduct) al controlador.
func NewCreateProductController (uc *application.CreateProduct) *CreateProductController {
	return &CreateProductController{uc: uc}
}

func (ctrl *CreateProductController) Run(c *gin.Context){
	var products domain.Product

	if err := c.ShouldBindJSON(&products); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "todos los campos tienen que ser requeridos"})
		return 
	}

	 err := ctrl.uc.Run(products) 

	if err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	return
	} else {
		c.JSON(http.StatusCreated, gin.H {"mensaje": "Producto creado"})
		c.JSON(http.StatusOK, products)
	}
}


