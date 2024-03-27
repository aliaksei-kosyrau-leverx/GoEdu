package application

import "time"

type UserRegistrationView struct {
	Id               string
	FirstName        string
	LastName         string
	Email            string
	Status           string
	RegistrationDate time.Time
}
