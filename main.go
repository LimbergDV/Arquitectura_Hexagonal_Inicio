package main

import (
	//alias para asignar a los recursos
	products "introduccion_go/src/products/infrastucture"
	routesProducts "introduccion_go/src/products/infrastucture/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	products.GoMySQL()
	r := gin.Default()
	routesProducts.Routes(r)
	r.Run()
}
