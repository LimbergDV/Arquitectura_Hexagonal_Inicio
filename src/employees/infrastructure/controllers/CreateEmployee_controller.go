package controllers

import (
	"introduccion_go/src/employees/application"
	"introduccion_go/src/employees/domain"
	"introduccion_go/src/employees/infrastructure"
	"introduccion_go/src/employees/infrastructure/routes/validators"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateEmployeeController struct {
	app *application.CreateEmployee
}

func NewCreateEmployeeController() *CreateEmployeeController {
	mysql := infrastructure.GetMySQL()
	app := application.NewCreateEmployee(mysql)
	return &CreateEmployeeController{app: app}
}

func (ce_c *CreateEmployeeController) Run (c *gin.Context){
	var employees domain.Employee

	if err := c.ShouldBindJSON(&employees); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": "Datos inválidos" + err.Error()})
		return
	}

	if err := validators.CheckEmployee(employees); err != nil {
		c.JSON(http.StatusBadRequest, gin.H {"status": false, "error": "Datos invalidos" + err.Error()})
	}
	
	rowsAffected, err := ce_c.app.Run(employees)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	if rowsAffected == 0{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, gin.H {"mensaje": "Producto creado"})
		c.JSON(http.StatusOK, employees)
	}
}