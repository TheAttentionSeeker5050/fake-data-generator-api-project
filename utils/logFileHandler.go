// A set of utility functions to use with the database
package utils

import (
	"fmt"
	"os"

	"log"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	file, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO - ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING - ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR - ", log.Ldate|log.Ltime|log.Lshortfile)
}

// This function writes errors in the log file
// it uses argument type CustomError struct
func WriteCustomError(customError CustomError) {
	// display error on console
	fmt.Errorf("%s: %s", customError.ErrorType, customError.Message)
	ErrorLogger.Println(fmt.Sprintf("%s: %s", customError.ErrorType, customError.Message))
}

func WriteCustomInfo(message string) {
	fmt.Println(message)
	InfoLogger.Println(fmt.Sprintf(message))
}
