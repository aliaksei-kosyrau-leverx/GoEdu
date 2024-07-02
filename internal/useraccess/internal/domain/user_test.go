package domain_test

import (
	"testing"

	"GoEdu/internal/useraccess/internal/domain"

	"github.com/stretchr/testify/require"
)

func TestUserRegistration(t *testing.T) {
	tests := map[string]struct {
		userName domain.UserName
		password string
		email    domain.UserEmail
	}{
		"successful creation userRegistration": {
			userName: domain.MustNewUserName("A", "A"),
			password: "aaa",
			email:    domain.MustNewUserEmail("aaa@gmail.com"),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			u, err := domain.NewUser(
				tc.userName,
				tc.password,
				tc.email,
			)
			sh := u.GetUserSnapshot()
			require.Nil(t, err)
			require.Equal(t, tc.password, sh.Password)
			require.Equal(t, tc.userName.GetFullName(), sh.UserName.GetFullName())
			require.Equal(t, tc.email, sh.Email)
		})
	}
}
