package controllers

import (
	"introduccion_go/src/employees/application"
	"introduccion_go/src/employees/infrastructure"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteEmployeeByIdController struct {
	app *application.DeleteEmployee
}

func NewDeleteEmployeeByIdController() *DeleteEmployeeByIdController {
	mysql := infrastructure.GetMySQL()
	app := application.NewDeleteEmployee(mysql)
	return &DeleteEmployeeByIdController{app: app}
}

func (ctrl *DeleteEmployeeByIdController) Run(c *gin.Context) {
	id := c.Param("id")
	employeeId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de empleado inválido"})
		return
	}

	rowsAffected, _ := ctrl.app.Run(employeeId)

	if rowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar el empleado"})
		return
	}

	// Devolviendo el mensaje y el ID eliminado
	c.JSON(http.StatusOK, gin.H{"message": "Empleado eliminado exitosamente"})
}
