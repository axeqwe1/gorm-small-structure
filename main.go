package main

import (
	"coffee-shop/db"
	"coffee-shop/routes"
)

// @title Coffee Shop API
// @version 1.0
// @description API for Coffee Shop
// @host localhost:8080
// @BasePath /

func main() {
	db.ConnectDatabase()
	r := routes.SetupRouter()
	r.Run(":8080")
}
