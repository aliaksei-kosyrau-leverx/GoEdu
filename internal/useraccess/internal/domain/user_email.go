package domain

import "net/mail"

type UserEmail struct {
	value string
}

func NewUserEmail(value string) (UserEmail, error) {
	_, err := mail.ParseAddress(value)
	if err != nil {
		return UserEmail{}, err
	}
	return UserEmail{value: value}, nil
}

func MustNewUserEmail(value string) UserEmail {
	email, err := NewUserEmail(value)
	if err != nil {
		panic(err)
	}
	return email
}

func (u UserEmail) String() string {
	return u.value
}
