package controllers

import (
	"introduccion_go/src/products/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteProductByIdController struct{
	uc *application.DeleteProduct
}

func NewDeleteProductByIdController (uc *application.DeleteProduct) *DeleteProductByIdController {
	return &DeleteProductByIdController{uc: uc}
}

func (ctrl *DeleteProductByIdController) Run(c *gin.Context)  {
	id := c.Param("id")
	productId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id de producto invalido"})
		return
	}
	err = ctrl.uc.Run(productId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Devolviendo el mensaje y el id eliminado
	c.JSON(http.StatusOK, gin.H{"message": "Producto eliminado exitosamente"})
}