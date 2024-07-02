package domain_test

import (
	"testing"

	"GoEdu/internal/useraccess/internal/domain"

	"github.com/stretchr/testify/require"
)

func TestPositiveUserPassword(t *testing.T) {
	tests := map[string]struct {
		password string
	}{
		"correct password": {
			password: "aaaa",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			password, err := domain.NewUserPassword(tc.password)
			require.Nil(t, err)
			require.Equal(t, tc.password, password.String())
		})
	}
}

func TestNegativeUserPassword(t *testing.T) {
	tests := map[string]struct {
		password string
	}{
		"incorrect password": {
			password: "",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := domain.NewUserPassword(tc.password)
			require.NotNil(t, err)
		})
	}
}

func TestPositiveHashedUserPassword(t *testing.T) {
	tests := map[string]struct {
		password domain.UserPassword
	}{
		"correct password": {
			password: domain.MustNewUserPassword("aaa"),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			password := domain.NewHashedUserPassword(tc.password)
			require.Equal(t, tc.password.String(), password.String())
		})
	}
}

func TestPositiveMustUserPassword(t *testing.T) {
	tests := map[string]struct {
		password string
	}{
		"correct password": {
			password: "aaaa",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			password := domain.MustNewUserPassword(tc.password)
			require.Equal(t, tc.password, password.String())
		})
	}
}

func TestNegativeMustUserPassword(t *testing.T) {
	tests := map[string]struct {
		password string
	}{
		"incorrect password": {
			password: "",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			require.Panics(t, func() {
				domain.MustNewUserPassword(tc.password)
			})
		})
	}
}
