package dependencies

import (
	"fmt"
	core "introduccion_go/src/Core"
	"introduccion_go/src/products/infrastucture"
)

var (mySQL infraestructure.MySQL)

func Init () {
	db, err := core.Conn_MySQL()
	if err != nil {
		fmt.Println("Error de servidor")
		return
	}

	mySQL = *infraestructure.NewMySQL(db)
}