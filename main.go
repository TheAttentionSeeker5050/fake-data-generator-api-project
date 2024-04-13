// This is the main entry point for the application, it will load the routes and start the server
// It will also load the configuration files, as well as the templates and static files into the gin router
package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"

	"example.com/main/routes"
)

func main() {

	router := gin.Default()

	// load templates
	router.LoadHTMLGlob("templates/**/*")

	// load static files
	router.Static("/assets", "./assets")

	routes.MainRoutes(router)
	routes.EndpointRoutes(router)

	err := router.Run() // listen and serve on 0.0.0.0:8080

	// in case an error occurs, save to file logs.log and re run the application
	if err != nil {
		file, fileErr := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if fileErr != nil {
			log.Fatal(fileErr)
		}
		log.SetOutput(file)
		log.Error("Application crashed: ", err.Error())

		// save the file
		file.Close()
	}

}
