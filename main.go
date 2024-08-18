package main

import (
	"fmt"

	"github.com/Jane-Mwangi/GoEventApi/db"
	"github.com/Jane-Mwangi/GoEventApi/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	fmt.Println("Server is starting on port 8080...")

	server.Run(":8080")
}
