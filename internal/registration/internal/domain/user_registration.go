package domain

import (
	"fmt"
	"github.com/google/uuid"
	"net/mail"
	"time"
)

var (
	UserRegistrationCannotBeConfirmedMoreThanOnceError = fmt.Errorf("UserRegistration can't be confirmed more than once")
)

func RegisterNewUser(firstName, lastName, email, password string) (*UserRegistration, error) {
	userMail, err := CreateUserEmail(email)
	if err != nil {
		return nil, err
	}

	return &UserRegistration{
		ID:               UserRegistrationID{Value: uuid.New()},
		Status:           WaitForConfirmation,
		Email:            userMail,
		Name:             firstName + " " + lastName,
		FirstName:        firstName,
		LastName:         lastName,
		Password:         password,
		RegistrationDate: time.Now(),
		ConfirmationDate: time.Time{},
	}, nil
}

type UserRegistration struct {
	ID               UserRegistrationID
	Status           UserRegistrationStatus
	Email            UserRegistrationEmail
	Name             string
	FirstName        string
	LastName         string
	Password         string
	RegistrationDate time.Time
	ConfirmationDate time.Time
}

func (u *UserRegistration) Confirm() error {
	if u.Status == Confirmed {
		return UserRegistrationCannotBeConfirmedMoreThanOnceError
	}
	u.Status = Confirmed
	u.ConfirmationDate = time.Now()
	return nil
}

type UserRegistrationEvent struct {
	ID               UserRegistrationID
	email            UserRegistrationEmail
	Name             string
	FirstName        string
	LastName         string
	RegistrationDate time.Time
}

type UserRegistrationID struct {
	Value uuid.UUID
}

func CreateUserEmail(value string) (UserRegistrationEmail, error) {
	_, err := mail.ParseAddress(value)
	if err != nil {
		return UserRegistrationEmail{}, err
	}
	return UserRegistrationEmail{value: value}, nil
}

type UserRegistrationEmail struct {
	value string
}

type UserRegistrationStatus string

const (
	WaitForConfirmation UserRegistrationStatus = "wait"
	Confirmed           UserRegistrationStatus = "confirmed"
)
