package controllers

import (
	"introduccion_go/src/products/application"
	infraestructure "introduccion_go/src/products/infrastucture"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteProductByIdController struct{
	app *application.DeleteProduct
}

func NewDeleteProductByIdController() *DeleteProductByIdController {
	mysql := infraestructure.GetMySQL()
	app := application.NewDeleteProduct(mysql)
	return &DeleteProductByIdController{app: app}
}

func (ctrl *DeleteProductByIdController) Run(c *gin.Context)  {
	id := c.Param("id")
	productId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id de producto invalido"})
		return
	}
	rowsAffected, _ := ctrl.app.Run(productId)

	if rowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar"})
		return
	}

	// Devolviendo el mensaje y el id eliminado
	c.JSON(http.StatusOK, gin.H{"message": "Producto eliminado exitosamente"})
}