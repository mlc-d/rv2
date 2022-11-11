package internal

import "time"

type User struct {
	id        *uint8
	name      string
	deletedAt time.Time
}

func (u *User) GetID() *uint8 {
	return u.id
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) IsDeleted() time.Time {
	return u.deletedAt
}

func NewUser(name string) *User {
	return &User{
		name: name,
	}
}
