package main

import (
	"github.com/gin-gonic/gin"

	// import main.routes.go
	"example.com/main/routes"
)

func main() {
	router := gin.Default()

	// load templates
	router.LoadHTMLGlob("templates/**/*")

	// load static files
	router.Static("/assets", "./assets")

	// import main.routes.go
	routes.MainRoutes(router)
	routes.EndpointRoutes(router)

	router.Run() // listen and serve on 0.0.0.0:8080
}
