package main

import (
	"introduccion_go/src/products/infrastucture/dependencies"
	"introduccion_go/src/products/infrastucture/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	dependencies.Init()
	r := gin.Default()
	routes.Routes(r)
	r.Run()
}
