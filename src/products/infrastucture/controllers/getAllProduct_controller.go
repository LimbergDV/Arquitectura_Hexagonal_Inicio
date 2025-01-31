package controllers

import (
	"introduccion_go/src/products/application"
	"net/http"
	

	"github.com/gin-gonic/gin"
)

type GetAllProductController struct {
	uc *application.GetAllProduct
}

func NewGetAllProductController (uc *application.GetAllProduct) *GetAllProductController{
	return &GetAllProductController{uc: uc}
}

func (ctrl *GetAllProductController) Run (c *gin.Context) {
	products, err := ctrl.uc.Run()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"products": products})
}