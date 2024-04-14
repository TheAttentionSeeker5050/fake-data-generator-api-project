// This is the main entry point for the application, it will load the routes and start the server
// It will also load the configuration files, as well as the templates and static files into the gin router
package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"example.com/main/config"
	"example.com/main/routes"
	"example.com/main/utils"
	"github.com/joho/godotenv"
)

// The main function
func main() {

	// Load the .env file in the current directory
	godotenv.Load()

	router := gin.Default()

	if os.Getenv("GIN_MODE") != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}

	// connect to database
	dbError := config.ConnectDatabase()
	// in case an error occurs, save to file logs.log and re run the application
	if dbError != nil {
		customError := utils.CustomError{
			Message:   dbError.Error(),
			ErrorType: utils.DB_ERROR,
		}

		utils.WriteCustomError(customError)

		// paning because there is no point in running this program without the databases
		panic(dbError)
	}

	// load templates
	router.LoadHTMLGlob("templates/**/*")

	// load static files
	router.Static("/assets", "./assets")

	routes.MainRoutes(router)
	routes.EndpointRoutes(router)

	err := router.Run() // listen and serve on 0.0.0.0:8080

	// in case an error occurs, save to file logs.log and re run the application
	if err != nil {
		customError := utils.CustomError{
			Message:   dbError.Error(),
			ErrorType: utils.SERVER_ERROR,
		}

		utils.WriteCustomError(customError)
	}

}
