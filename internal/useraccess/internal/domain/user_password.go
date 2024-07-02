package domain

import (
	"errors"
	"strings"
)

var ErrPasswordCannotBeEmpty = errors.New("password is empty")

type UserPassword struct {
	value string
}

func NewUserPassword(value string) (UserPassword, error) {
	if len(strings.TrimSpace(value)) == 0 {
		return UserPassword{}, ErrPasswordCannotBeEmpty
	}
	return UserPassword{value}, nil
}

func MustNewUserPassword(value string) UserPassword {
	password, err := NewUserPassword(value)
	if err != nil {
		panic(err)
	}
	return password
}

func (u UserPassword) String() string {
	return u.value
}

type HashedUserPassword struct {
	value string
}

// For educational purposes we use different types between UserPassword input VO and
// HashedPassword to explicitly mark that password for domain model is encrypted.
func NewHashedUserPassword(password UserPassword) HashedUserPassword {
	return HashedUserPassword{value: password.String()}
}

func (h HashedUserPassword) String() string {
	return h.value
}
