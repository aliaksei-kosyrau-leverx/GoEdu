package domain

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserID(t *testing.T) {
	tests := map[string]struct {
		id string
	}{
		"case 1": {
			id: "12345678123456781234567812345678",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			urid := MustParseUserRegistrationID(tc.id)
			require.NotEmpty(t, urid)
			require.Equal(t, tc.id, strings.Replace(urid.String(), "-", "", -1))
		})
	}
}