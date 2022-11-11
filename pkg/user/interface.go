package user

import (
	"time"

	"gitlab.com/mlc-d/rv2/pkg/user/internal"
)

type UserService interface {
	Create(name string) *User // TODO: handle error
	Delete(id *uint8) error
}

// userService implements the correspondent interface
type userService struct {
	repo internal.UserRepo
}

func (us *userService) Create(name string) *User {
	u := internal.NewUser(name)
	id, err := us.repo.Insert(u)
	if err != nil {
		return nil
	}

	return &User{
		ID:   id,
		Name: name,
	}
}

func (us *userService) Delete(id *uint8) *User {
	return nil
}

type User struct {
	ID        *int64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}
