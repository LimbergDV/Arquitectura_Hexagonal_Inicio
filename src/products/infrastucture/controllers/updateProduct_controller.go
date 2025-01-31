package controllers

import (
	"fmt"
	"introduccion_go/src/products/application"
	"introduccion_go/src/products/domain"
	infraestructure "introduccion_go/src/products/infrastucture"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateProductByIdController struct{
	app *application.UpdateProduct
}

func NewUpdateProductByIdController() *UpdateProductByIdController{
	mysql := infraestructure.GetMySQL()
	app := application.NewUpdateProduct(mysql)
	return &UpdateProductByIdController{app: app}
}

func (ctrl *UpdateProductByIdController) Run(c *gin.Context){
	id := c.Param("id")
	var product domain.Product

	id_product, _ := strconv.ParseUint(id, 10, 64)

	if err :=  c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	RowsAffected, _ := ctrl.app.Run(int(id_product), product)

	if RowsAffected == 0{
		fmt.Print("hola")
	}

	// Send a successful response
	c.JSON(http.StatusOK, gin.H{
		"message":      "Product updated successfully",
		"id":           idInt,
	})

}