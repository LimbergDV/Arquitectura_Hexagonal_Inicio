package routes

import (
	"introduccion_go/src/employees/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func Routes (r *gin.Engine) {
	
	employeesRoutes := r.Group("/employees") 
	{
		employeesRoutes.POST("/", controllers.NewCreateEmployeeController().Run)
		employeesRoutes.GET("/", controllers.NewGetAllEmployeesController().Run)
		employeesRoutes.PUT("/:id", controllers.NewUpdateEmployeeByIdController().Run)
		employeesRoutes.DELETE("/:id", controllers.NewDeleteEmployeeByIdController().Run)
	}
}