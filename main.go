package main

import (
	"github.com/flpcastro/apirest-go-gin/database"
	"github.com/flpcastro/apirest-go-gin/routes"
)

func main() {
	database.DatabaseConnect()
	routes.HandleRequest()
}
