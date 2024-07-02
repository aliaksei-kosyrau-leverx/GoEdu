package domain_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"GoEdu/internal/useraccess/internal/domain"
)

func TestPositiveNewUserEmail(t *testing.T) {
	tests := map[string]struct {
		email string
	}{
		"successful create new email": {
			email: "email@gmail.com",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			email, err := domain.NewUserEmail(tc.email)
			require.Nil(t, err)
			require.Equal(t, tc.email, email.String())
		})
	}
}

func TestNegativeNewUserEmail(t *testing.T) {
	tests := map[string]struct {
		email string
	}{
		"unsuccessful create new email with empty string": {
			email: "",
		},
		"unsuccessful create new email in incorrect way": {
			email: "@.com",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := domain.NewUserEmail(tc.email)
			require.NotNil(t, err)
		})
	}
}

func TestPositiveMustNewUserEmail(t *testing.T) {
	tests := map[string]struct {
		email string
	}{
		"successful create new email": {
			email: "email@gmail.com",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			email := domain.MustNewUserEmail(tc.email)
			require.Equal(t, tc.email, email.String())
		})
	}
}

func TestNegativeMustNewUserEmail(t *testing.T) {
	tests := map[string]struct {
		email string
	}{
		"unsuccessful create new email in incorrect way": {
			email: "@.com",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			require.Panics(t, func() {
				domain.MustNewUserEmail(tc.email)
			})
		})
	}
}
