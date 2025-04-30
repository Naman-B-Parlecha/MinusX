package main

import (
	"fmt"

	"github.com/Naman-B-Parlecha/MinusX/internal/db"
	"github.com/Naman-B-Parlecha/MinusX/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := db.ConnectDb()
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}

	fmt.Println("Connected to the database")
	r := gin.Default()

	routes.SetupRoutes(r, db)

	fmt.Println("Starting server on port 8081")
	r.Run(":8081")
}
