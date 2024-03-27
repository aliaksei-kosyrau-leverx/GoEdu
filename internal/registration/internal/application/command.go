package application

import "GoEdu/internal/registration/internal/domain"

type Command interface {
}

type CommandResult interface {
}

type RegisterNewUserCommand struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type ConfirmUserRegistrationCommand struct {
	ID domain.UserRegistrationID
}
