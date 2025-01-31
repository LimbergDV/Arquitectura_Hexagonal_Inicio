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
	res := ctrl.app.Run()

	if len(res) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": false, "error": "No se encontró ningún producto"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"products": res})
	}

	
}