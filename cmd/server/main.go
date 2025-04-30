package main

import (
	"github.com/Naman-B-Parlecha/MinusX/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":8080")
}
