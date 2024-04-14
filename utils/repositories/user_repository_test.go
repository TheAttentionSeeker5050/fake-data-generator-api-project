package repositories

import (
	"os"
	"testing"
	"time"

	"example.com/main/config"
	"example.com/main/models"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestUsersRepository(t *testing.T) {
	if os.Getenv("GIN_MODE") != "GITHUB_ACTIONS" {

		// Load the .env file in the current directory
		envVariableErr := godotenv.Load("../../.env")
		assert.Nil(t, envVariableErr, "Expected no error loading .env file")
		if envVariableErr != nil {
			return
		}
	}

	// Get the connection uri from the environment variables
	connUri := os.Getenv("MYSQL_TEST_DB_URI")

	// Connect to the test database
	connErr := config.ConnectGormDatabase(connUri)
	assert.Nil(t, connErr, "Expected no error connecting to the database")
	if connErr != nil {
		return
	}

	// Get the DB instance from the config package
	db := config.DB
	assert.NotNil(t, db, "Expected DB instance to be created")
	if db == nil {
		return
	}

	// Create a new user repository
	repo := NewGormUserRepository(db)

	// Assert that the repository is created with the correct values
	assert.NotNil(t, repo, "Expected repository to be created")
	assert.Equal(t, db, repo.Conn, "Expected connection to match the DB instance")
	if repo == nil {
		return
	}

	// Create a new user model
	user := &models.UserModel{
		Username:  "testuser",
		FirstName: "Test",
		LastName:  "User",
		Email:     "user@email.com",
		Password:  "password",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Create the user in the database
	createErr := repo.CreateUser(user)
	assert.Nil(t, createErr, "Expected no error creating user")
	if createErr != nil {
		return
	}

	// Get the user by its ID
	userOnDb, getErr := repo.GetUserById(user.ID)
	assert.Nil(t, getErr, "Expected no error getting user by ID")
	assert.Equal(t, user.ID, userOnDb.ID, "Expected user ID to match")
	if getErr != nil {
		return
	}

	// Get the password of the user by its username
	password, getPassErr := repo.GetPasswordByUsername(user.Username)
	assert.Nil(t, getPassErr, "Expected no error getting password by username")
	assert.Equal(t, user.Password, password, "Expected password to match")
	if getPassErr != nil {
		return
	}

	// Delete the user from the database
	deleteErr := db.Delete(user)
	assert.Nil(t, deleteErr.Error, "Expected no error deleting user")
	if deleteErr.Error != nil {
		return
	}
}
