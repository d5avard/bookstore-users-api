package users

import (
	"strings"

	"github.com/d5avard/bookstore-users-api/utils/errors"
)

type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first-name"`
	LastName    string `json:"last-name"`
	Email       string `json:"email"`
	DateCreated string `json:"date-created"`
}

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	return nil
}
