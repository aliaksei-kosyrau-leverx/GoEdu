package application_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"GoEdu/internal/useraccess/internal/application"
	"GoEdu/internal/useraccess/internal/domain"
)

var _ application.UserRepository = (*CreateNewUserRepoMock)(nil)

type CreateNewUserRepoMock struct {
	added bool
}

func (r *CreateNewUserRepoMock) Add(_ context.Context, _ *domain.User) error {
	r.added = true
	return nil
}
func (r CreateNewUserRepoMock) Load(
	_ context.Context,
	_ domain.UserEmail) (*domain.User, error) {
	return &domain.User{}, nil
}
func (r CreateNewUserRepoMock) Update(_ context.Context, _ *domain.User) error {
	return nil
}

var _ application.PasswordManager = (*PasswordHasherSpyChecker)(nil)

type PasswordHasherSpyChecker struct {
	hashed bool
}

func (p *PasswordHasherSpyChecker) IsEqual(_ domain.UserPassword, _ domain.HashedUserPassword) bool {
	p.hashed = true
	return true
}

func TestPositiveCreateNewUserRegistration(t *testing.T) {
	tests := map[string]struct {
		repository *CreateNewUserRepoMock
		command    application.CreateUserCommand
	}{
		"successful Create user registration": {
			repository: &CreateNewUserRepoMock{},
			command: application.CreateUserCommand{
				FirstName: "a",
				LastName:  "a",
				Email:     domain.MustNewUserEmail("email@mail.ru"),
				Password:  "password",
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			handler := application.NewCreateUserCommandHandler(tc.repository)
			require.NotNil(t, handler, "handler is nil")
			_, err := handler.Handle(context.Background(), tc.command)
			require.Nil(t, err, "error not nil")
			require.Equal(t, true, tc.repository.added, "user registration not added")
		})
	}
}
