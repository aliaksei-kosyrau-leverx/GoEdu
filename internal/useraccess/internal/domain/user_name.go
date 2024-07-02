package domain

import (
	"errors"
	"strings"
)

var ErrFirstNameCannotBeEmpty = errors.New("first name is empty")
var ErrLastNameCannotBeEmpty = errors.New("last name is empty")

type UserName struct {
	firstName string
	lastName  string
}

func NewUserName(firstName, lastName string) (UserName, error) {
	if len(strings.TrimSpace(firstName)) == 0 {
		return UserName{}, ErrFirstNameCannotBeEmpty
	}
	if len(strings.TrimSpace(lastName)) == 0 {
		return UserName{}, ErrLastNameCannotBeEmpty
	}

	return UserName{firstName: firstName, lastName: lastName}, nil
}

func MustNewUserName(firstName, lastName string) UserName {
	userName, err := NewUserName(firstName, lastName)
	if err != nil {
		panic(err)
	}
	return userName
}

func (u UserName) GetFirstName() string {
	return u.firstName
}

func (u UserName) GetLastName() string {
	return u.lastName
}

func (u UserName) GetFullName() string {
	return u.firstName + " " + u.lastName
}
