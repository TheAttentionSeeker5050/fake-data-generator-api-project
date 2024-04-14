// This package contains the configuration methods for the project
package config

// import mysql driver
import (
	"fmt"
	"os"

	// database models
	models "example.com/main/models"
	"example.com/main/utils"

	// database driver imports
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB connects to the database
func ConnectDatabase() error {

	// make the address with environment variables
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_DB_USER"),
		os.Getenv("MYSQL_DB_PASSWORD"),
		os.Getenv("MYSQL_DB_HOST"),
		os.Getenv("MYSQL_DB_PORT"),
		os.Getenv("MYSQL_DB_NAME"),
	)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return fmt.Errorf("Failed to connect to database: \n%s", err.Error())
	}

	err = database.AutoMigrate(&models.UserModel{})
	if err != nil {
		// panic(fmt.Sprintf("Failed to migrate database: \n%s", err.Error()))
		return fmt.Errorf("Failed to migrate database: \n%s", err.Error())
	}

	DB = database

	// when successfully initiating the database and migrating models, return a message log
	utils.WriteCustomInfo("Successfully initiated the database")

	return nil
}
