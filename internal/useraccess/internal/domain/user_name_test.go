package domain_test

import (
	"testing"

	"GoEdu/internal/useraccess/internal/domain"

	"github.com/stretchr/testify/require"
)

func TestPositiveUserName(t *testing.T) {
	tests := map[string]struct {
		firstName string
		lastName  string
	}{
		"correct name example": {
			"jotaro",
			"joske",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			uName, err := domain.NewUserName(tc.firstName, tc.lastName)
			require.Nil(t, err)
			require.Equal(t, tc.firstName, uName.GetFirstName())
			require.Equal(t, tc.lastName, uName.GetLastName())
			require.Equal(t, tc.firstName+" "+tc.lastName, uName.GetFullName())
		})
	}
}

func TestNegativeUserName(t *testing.T) {
	tests := map[string]struct {
		firstName string
		lastName  string
	}{
		"incorrect first name example": {
			firstName: "",
			lastName:  "B",
		},
		"incorrect last name example": {
			firstName: "A",
			lastName:  "",
		},
		"incorrect first and last name example": {
			firstName: " ",
			lastName:  " ",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := domain.NewUserName(tc.firstName, tc.lastName)
			require.NotNil(t, err)
		})
	}
}

func TestPositiveMustUserName(t *testing.T) {
	tests := map[string]struct {
		firstName string
		lastName  string
	}{
		"correct name example": {
			"jotaro",
			"joske",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			uName := domain.MustNewUserName(tc.firstName, tc.lastName)
			require.Equal(t, tc.firstName, uName.GetFirstName())
			require.Equal(t, tc.lastName, uName.GetLastName())
			require.Equal(t, tc.firstName+" "+tc.lastName, uName.GetFullName())
		})
	}
}

func TestNegativeMustUserName(t *testing.T) {
	tests := map[string]struct {
		firstName string
		lastName  string
	}{
		"incorrect first and last name example": {
			firstName: " ",
			lastName:  " ",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			require.Panics(t, func() {
				domain.MustNewUserName(tc.firstName, tc.lastName)
			})
		})
	}
}
