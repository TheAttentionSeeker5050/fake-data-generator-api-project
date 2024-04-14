// This package contains the gorm and mongo transactions repositories for our models
package repositories

import (
	"fmt"

	"example.com/main/models"
	"gorm.io/gorm"
)

// A gorm repository for user model
type GormUserRepository struct {
	Conn *gorm.DB
}

// This function creates an object that represents the gorm user repository
func NewGormUserRepository(conn *gorm.DB) *GormUserRepository {
	return &GormUserRepository{Conn: conn}
}

// This function creates a new user in the database using the gorm repository
func (r *GormUserRepository) CreateUser(user *models.UserModel) error {

	// make sure the username and email are unique
	userOnDb := &models.UserModel{}
	r.Conn.Where("username = ?", user.Username).Or("email = ?", user.Email).First(userOnDb)

	// if the user already exists, return an error
	if userOnDb.ID != 0 {
		return fmt.Errorf("username or email already exists")
	}

	// create the user
	result := r.Conn.Create(user)

	// if there is an error, return it
	return result.Error
}

// This function returns a user by its id using the gorm repository
func (r *GormUserRepository) GetUserById(id uint) (*models.UserModel, error) {
	// create a new empty user model object
	user := &models.UserModel{}

	// return everything except the password, save it in the user object
	result := r.Conn.Select("id, username, first_name, last_name, email, created_at, updated_at").First(user, id)

	// return the user and the error
	return user, result.Error
}

// This function returns the password of a user by its username using the gorm repository
func (r *GormUserRepository) GetPasswordByUsername(username string) (string, error) {
	// create a new empty user model object
	user := &models.UserModel{}

	// return the password, save it in the user object
	result := r.Conn.Select("password").Where("username = ?", username).First(user)

	// return the password and the error
	return user.Password, result.Error
}

// This function changes the password of a user by its username using the gorm repository. It will store the new password as it is in the parameter
func (r *GormUserRepository) ChangePasswordByUsername(username string, newPassword string) error {
	// update the password
	result := r.Conn.Model(&models.UserModel{}).Where("username = ?", username).Update("password", newPassword)

	// return the error
	return result.Error
}
