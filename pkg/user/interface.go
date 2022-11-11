package user

import "gitlab.com/mlc-d/rv2/pkg/user/internal"

type UserService interface {
	Create(name string) *User
	Delete(id *uint8) error
}

// userService implements the correspondent interface
type userService struct {
	repo internal.UserRepo
}

type User struct {
	Id        *uint8    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}
