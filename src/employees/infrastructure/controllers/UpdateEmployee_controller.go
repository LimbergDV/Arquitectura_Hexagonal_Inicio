package controllers

import (
	"fmt"
	"introduccion_go/src/employees/application"
	"introduccion_go/src/employees/domain"
	"introduccion_go/src/employees/infrastructure"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateEmployeeByIdController struct {
	app *application.UpdateEmployee
}

func NewUpdateEmployeeByIdController() *UpdateEmployeeByIdController {
	mysql := infrastructure.GetMySQL()
	app := application.NewUpdateEmployee(mysql)
	return &UpdateEmployeeByIdController{app: app}
}

func (ctrl *UpdateEmployeeByIdController) Run(c *gin.Context) {
	id := c.Param("id")
	var employee domain.Employee

	idEmployee, _ := strconv.ParseUint(id, 10, 64)

	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, _ := ctrl.app.Run(int(idEmployee), employee)

	if rowsAffected == 0 {
		fmt.Println("No se pudo actualizar el empleado")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar el empleado"})
		return
	}

	// Enviar una respuesta exitosa
	c.JSON(http.StatusOK, gin.H{
		"message": "Employee updated successfully",
	})
}
