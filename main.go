package main

import (
	"gitlab.com/alura-courses-code/golang/go-gin-rest-api/database"
	"gitlab.com/alura-courses-code/golang/go-gin-rest-api/routes"
)

func main() {
	database.ConnectDatabase()
	routes.HandleRequests()
}
