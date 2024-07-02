package domain

import "time"

type UserEvent struct {
	ID               string
	Email            string
	Name             string
	FirstName        string
	LastName         string
	RegistrationDate time.Time
}
