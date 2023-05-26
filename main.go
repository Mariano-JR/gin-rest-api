package main

import (
	"github.com/Mariano-JR/gin-rest-api/database"
	"github.com/Mariano-JR/gin-rest-api/routes"
)

func main() {
	database.ConectaDB()
	routes.HandleRequests()
}
