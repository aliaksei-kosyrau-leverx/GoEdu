package domain

import "github.com/google/uuid"

type UserID uuid.UUID

func NewUserID() UserID {
	return UserID(uuid.New())
}

func (u UserID) String() string {
	return uuid.UUID(u).String()
}

func ParseUserID(urid string) (UserID, error) {
	uid, err := uuid.Parse(urid)
	if err != nil {
		return UserID{}, err
	}
	return UserID(uid), nil
}

func MustParseUserID(urid string) UserID {
	uid, err := ParseUserID(urid)
	if err != nil {
		panic(err)
	}
	return uid
}
