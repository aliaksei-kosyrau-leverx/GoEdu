package application

import (
	"context"

	"GoEdu/internal/useraccess/internal/domain"
)

type CreateUserCommand struct {
	FirstName string
	LastName  string
	Email     domain.UserEmail
	Password  string
}

type CreateUserCommandResult struct {
}

func NewCreateUserCommandHandler(
	repository UserRepository) *CreateUserCommandHandler {
	return &CreateUserCommandHandler{repository: repository}
}

type CreateUserCommandHandler struct {
	repository UserRepository
}

func (r *CreateUserCommandHandler) Handle(ctx context.Context, command Command) (CommandResult, error) {
	createUserCommand := command.(CreateUserCommand)

	hashedPassword := createUserCommand.Password

	userName, err := domain.NewUserName(createUserCommand.FirstName, createUserCommand.LastName)
	if err != nil {
		return CreateUserCommandResult{}, err
	}

	user, err := domain.NewUser(
		userName,
		hashedPassword,
		createUserCommand.Email,
	)
	if err != nil {
		return CreateUserCommandResult{}, err
	}

	err = r.repository.Add(ctx, user)
	if err != nil {
		return CreateUserCommandResult{}, err
	}

	return CreateUserCommandResult{}, nil
}
