package services

import (
	"github.com/d5avard/bookstore-users-api/domain/users"
	"github.com/d5avard/bookstore-users-api/utils/errors"
)

func CreateUser(user *users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return user, nil
}

func GetUser(ID int64) (*users.User, *errors.RestErr) {
	user := new(users.User)
	if err := user.Get(ID); err != nil {
		return nil, err
	}
	return user, nil
}
