package controllers

import (
	"introduccion_go/src/products/application"
	"introduccion_go/src/products/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateProductByIdController struct{
	uc *application.UpdateProduct
}

func NewUpdateProductByIdController(uc *application.UpdateProduct) *UpdateProductByIdController{
	return &UpdateProductByIdController{uc: uc}
}

func (ctrl *UpdateProductByIdController) Run(c *gin.Context){
	var product domain.Product

	if err :=  c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	
	if err := ctrl.uc.Run(idInt, product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Send a successful response
	c.JSON(http.StatusOK, gin.H{
		"message":      "Product updated successfully",
		"id":           idInt,
	})

}