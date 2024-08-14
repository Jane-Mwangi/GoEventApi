package main

import (
	"github.com/Jane-Mwangi/GoEventApi/db"
	"github.com/Jane-Mwangi/GoEventApi/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	

	server.Run(":8080")
}


