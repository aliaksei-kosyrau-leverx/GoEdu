package domain

import (
	"errors"
)

var (
	ErrUserRegistrationCannotBeConfirmedMoreThanOnce = errors.New("user can't be confirmed more than once")
)

type User struct {
	id       UserID
	email    UserEmail
	userName UserName
	password HashedUserPassword
}

func NewUser(
	userName UserName,
	password string,
	email UserEmail,
) (*User, error) {
	return &User{
		NewUserID(),
		email,
		userName,
		NewHashedUserPassword(MustNewUserPassword(password)),
	}, nil
}

func (u *User) GetUserSnapshot() *UserSnapshot {
	return &UserSnapshot{
		ID:       u.id,
		Email:    u.email,
		UserName: u.userName,
		Password: u.password,
	}
}

type UserSnapshot struct {
	ID       UserID
	Email    UserEmail
	UserName UserName
	Password HashedUserPassword
}

// func MustCreateUserRegistrationFromSnapshot(snapshot UserRegistrationSnapshot) *UserRegistration {
// 	return &UserRegistration{
// 		snapshot.ID,
// 		snapshot.Status,
// 		snapshot.Email,
// 		snapshot.UserName,
// 		snapshot.Password,
// 		snapshot.RegistrationDate,
// 		snapshot.ConfirmationDate,
// 	}
// }

// func (u *UserRegistration) GetSnapshot() UserRegistrationSnapshot {
// 	return MustCreateUserRegistrationSnapshot(
// 		u.id,
// 		u.status,
// 		u.email,
// 		u.userName,
// 		u.password,
// 		u.registrationDate,
// 		u.confirmationDate,
// 	)
// }

// type UserRegistrationSnapshot struct {
// 	ID               UserID
// 	Status           UserRegistrationStatus
// 	Email            UserEmail
// 	UserName         UserName
// 	Password         HashedUserPassword
// 	RegistrationDate time.Time
// 	ConfirmationDate time.Time
// }

// func MustCreateUserRegistrationSnapshot(
// 	id UserID,
// 	status UserRegistrationStatus,
// 	email UserEmail,
// 	name UserName,
// 	password HashedUserPassword,
// 	registrationDate time.Time,
// 	confirmationDate time.Time,

// ) UserRegistrationSnapshot {
// 	return UserRegistrationSnapshot{
// 		id,
// 		status,
// 		email,
// 		name,
// 		password,
// 		registrationDate,
// 		confirmationDate,
// 	}
// }
