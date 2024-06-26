// This package contains the gorm and mongo models
package models

import "time"

// User model for Gorm ORM using MySQL database
type UserModel struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Username  string    `gorm:"unique" json:"username"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
