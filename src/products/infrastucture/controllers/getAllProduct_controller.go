package controllers

import (
	"introduccion_go/src/products/application"
	infraestructure "introduccion_go/src/products/infrastucture"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllProductController struct {
	app *application.GetAllProduct
}

func NewGetAllProductController () *GetAllProductController{
	mysql := infraestructure.GetMySQL()
	app := application.NewGetAllProduct(mysql)
	return &GetAllProductController{app: app}
}

func (ctrl *GetAllProductController) Run (c *gin.Context) {
	products := ctrl.app.Run()

	if products == nil {
		c.JSON(http.StatusNotFound, gin.H{"status": false, "error": "estas mal"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}