// This is the main entry point for the application, it will load the routes and start the server
// It will also load the configuration files, as well as the templates and static files into the gin router
package main

import (
	"context"
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
	envVariableErr := godotenv.Load()
	if envVariableErr != nil {
		utils.ErrorLogger.Println("Error loading .env file")
	}

	router := gin.Default()

	if os.Getenv("GIN_MODE") != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}

	// connect to database
	gormErr := config.ConnectGormDatabase()
	// in case an error occurs, save to file logs.log and re run the application
	if gormErr != nil {
		customError := utils.CustomError{
			Message:   gormErr.Error(),
			ErrorType: utils.DB_ERROR,
		}

		utils.WriteCustomError(customError)

		// paning because there is no point in running this program without the databases
		panic(gormErr)
	}

	// connect to mongo database
	mongoClient := config.ConnectMongoDatabase()

	// defer disconnecting from the database
	defer func() {
		// in case there is error disconnecting from the database, save to file logs.log and re run the application
		if mongoDbErr := mongoClient.Disconnect(context.TODO()); mongoDbErr != nil {
			customError := utils.CustomError{
				Message:   mongoDbErr.Error(),
				ErrorType: utils.DB_ERROR,
			}

			utils.WriteCustomError(customError)

			// paning because there is no point in running this program without the databases
			panic(mongoDbErr)
		}
	}()

	// load templates
	router.LoadHTMLGlob("templates/**/*")

	// load static files
	router.Static("/assets", "./assets")

	routes.MainRoutes(router)
	routes.EndpointRoutes(router)

	serverErr := router.Run() // listen and serve on 0.0.0.0:8080

	// in case an error occurs, save to file logs.log and re run the application
	if serverErr != nil {
		customError := utils.CustomError{
			Message:   serverErr.Error(),
			ErrorType: utils.SERVER_ERROR,
		}

		utils.WriteCustomError(customError)
	}

}
