package model

import (
	"errors"
	"time"
)

type User struct {
	ID        uint64    `json:"id"`
	Username  string    `json:"username"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Password  string    `json:"password"`
	DoB       time.Time `json:"dob" gorm:"column:dob"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

type UserMediaSocial struct {
	ID        uint64 `json:"id"`
	UserID    uint64 `json:"user_id"`
	Title     string `json:"title"`
	Url       string `json:"url"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

// https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples/
// https://gin-gonic.com/docs/examples/binding-and-validation/
type UserSignUp struct {
	Username  string `json:"username" binding:"required"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password" binding:"required"`
}

func (u UserSignUp) Validate() error {
	// check username
	if u.Username == "" {
		return errors.New("invalid username")
	}
	if len(u.Password) < 6 {
		return errors.New("invalid password")
	}
	return nil
}
