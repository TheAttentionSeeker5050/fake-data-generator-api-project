// This package contains the gorm models for the user entity
package models

type UserModel struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	Username  string `gorm:"unique" json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `gorm:"unique" json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
