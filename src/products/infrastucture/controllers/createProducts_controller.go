package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProduct (c *gin.Context){
	var Product struct{
		Name string `json: name`
		Price string `json: price`
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "false",
			"error": "Error 400, datos invalidos",
		})
		return
	}

	newProduct := product ProdcuProduct
	{
		Name:  product.Name,
		Price: product.Price,
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "false",
			"error":  "Error 500, no se puedo guardar el producto: " + err.Error(),
		})
		return
	}

	// Respuesta exitosa
	c.JSON(http.StatusCreated, gin.H{
		"status":  "true",
		"message": "Producto creado exitosamente",
		"product": newProduct,
	})
}


