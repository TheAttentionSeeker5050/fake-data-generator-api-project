// This package contains the configuration methods for the project
package config

// import mysql driver
import (
	"fmt"

	// database models
	models "example.com/main/models"
	"example.com/main/utils"

	// database driver imports
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Gorm db instance
var DB *gorm.DB

// Connects to the database using Gorm
func ConnectGormDatabase(dsn string) error {

	// open the database
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return fmt.Errorf("Failed to connect to database: \n%s", err.Error())
	}

	// migrate the database objects
	err = database.AutoMigrate(&models.UserModel{})
	if err != nil {
		return fmt.Errorf("Failed to migrate database: \n%s", err.Error())
	}

	// update the instance for use by the server utilities
	DB = database

	// when successfully initiating the database and migrating models, return a message log
	utils.WriteCustomInfo("Successfully initiated the database")

	return nil
}
