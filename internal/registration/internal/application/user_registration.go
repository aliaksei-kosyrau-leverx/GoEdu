package application

import "GoEdu/internal/registration/internal/domain"

type GetWaitingForConfirmationUserRegistrationsQuery struct {
}

type UserRegistrationReadModel interface {
	GetWaitingForConfirmationUserRegistrations(GetWaitingForConfirmationUserRegistrationsQuery) ([]domain.UserRegistration, error)
}
