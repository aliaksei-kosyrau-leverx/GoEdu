package application_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"GoEdu/internal/useraccess/internal/application"
	"GoEdu/internal/useraccess/internal/domain"
)

var _ application.LoginRegistrationRepository = (*LoginUserRepoMock)(nil)

type LoginUserRepoMock struct {
	added bool
}

func (r *LoginUserRepoMock) Add(_ context.Context, _ *domain.UserAuthentication) error {
	r.added = true
	return nil
}
func (r LoginUserRepoMock) Load(
	_ context.Context,
	_ domain.UserAuthenticationID) (*domain.UserAuthentication, error) {
	return &domain.UserAuthentication{}, nil
}
func (r LoginUserRepoMock) Update(_ context.Context, _ *domain.UserAuthentication) error {
	return nil
}

var _ application.UniqueEmailVerifier = (*UniqueEmailVerifierSpy)(nil)

type UniqueEmailVerifierSpy struct {
	checked bool
}

func (u *UniqueEmailVerifierSpy) IsUnique(_ context.Context, _ domain.UserAuthenticationEmail) error {
	u.checked = true
	return nil
}

var _ application.PasswordHasher = (*PasswordHasherSpy)(nil)

type PasswordHasherSpy struct {
	hashed bool
}

func (p *PasswordHasherSpy) Hash(password domain.UserPassword) (domain.HashedUserPassword, error) {
	p.hashed = true
	return domain.NewHashedUserPassword(password), nil
}

func TestPositiveLoginUserAuthentication(t *testing.T) {
	tests := map[string]struct {
		verifier   *UniqueEmailVerifierSpy
		hasher     *PasswordHasherSpy
		repository *LoginUserRepoMock
		command    application.LoginUserCommand
	}{
		"successful register user registration": {
			verifier:   &UniqueEmailVerifierSpy{},
			hasher:     &PasswordHasherSpy{},
			repository: &LoginUserRepoMock{},
			command: application.LoginUserCommand{
				FirstName: "a",
				LastName:  "a",
				Email:     domain.MustNewUserEmail("email@mail.ru"),
				Password:  domain.MustNewUserPassword("password"),
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			handler := application.NewLoginUserCommandHandler(tc.hasher, tc.repository, tc.verifier)
			require.NotNil(t, handler, "handler is nil")
			_, err := handler.Handle(context.Background(), tc.command)
			require.Nil(t, err, "error not nil")
			require.Equal(t, true, tc.verifier.checked, "verifier not checked")
			require.Equal(t, true, tc.hasher.hashed, "password not hashed")
			require.Equal(t, true, tc.repository.added, "user registration not added")
		})
	}
}