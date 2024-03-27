package application

import (
	"GoEdu/internal/registration/internal/domain"
	"context"
	"fmt"
)

type CommandHandler interface {
	Handle(context.Context, Command) (CommandResult, error)
}

type RegisterNewUserCommandHandler struct {
	hasher     PasswordHasher
	verifier   UniqueEmailVerifier
	repository UserRegistrationRepository
}

func (r *RegisterNewUserCommandHandler) Handle(context context.Context, command Command) (CommandResult, error) {
	regNewUserCommand := command.(RegisterNewUserCommand)

	h, err := r.hasher.Hash(regNewUserCommand.Password)
	if err != nil {
		return nil, err
	}

	user, err := domain.RegisterNewUser(
		regNewUserCommand.FirstName,
		regNewUserCommand.LastName,
		regNewUserCommand.Email,
		h,
	)
	if err != nil {
		return nil, err
	}

	if res, err := r.verifier.IsUnique(context, user.Email); err != nil || res != true {
		if err != nil {
			return nil, err
		} else {
			return nil, UserEmailMustBeUniqueError
		}
	}

	err = r.repository.Add(context, *user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

type ConfirmUserRegistrationCommandHandler struct {
	repository UserRegistrationRepository
}

func (c *ConfirmUserRegistrationCommandHandler) Handle(context context.Context, command Command) (CommandResult, error) {
	confUserRegistrationCommand := command.(ConfirmUserRegistrationCommand)
	user, err := c.repository.Load(context, confUserRegistrationCommand.ID)

	if err != nil {
		return nil, err
	}

	if err = user.Confirm(); err != nil {
		return nil, err
	}

	err = c.repository.Update(context, user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

var (
	UserEmailMustBeUniqueError = fmt.Errorf("User email must be unique")
)

type UniqueEmailVerifier interface {
	IsUnique(context.Context, domain.UserRegistrationEmail) (bool, error)
}

type PasswordHasher interface {
	Hash(string) (string, error)
}

type UserRegistrationRepository interface {
	Add(context.Context, domain.UserRegistration) error
	Load(context.Context, domain.UserRegistrationID) (domain.UserRegistration, error)
	Update(context.Context, domain.UserRegistration) error
}
