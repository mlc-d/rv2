package internal

import (
	"database/sql"

	"gitlab.com/mlc-d/rv2/database"
)

type UserRepo interface {
	Insert(u *User) (*int64, error)
	Delete(id *int64) error
}

func NewRepo() UserRepo {
	return &userRepo{
		db: database.GetDB(),
	}
}

type userRepo struct {
	db *sql.DB
}

func (ur *userRepo) Insert(u *User) (*int64, error) {
	return nil, nil
}

func (ur *userRepo) Delete(id *int64) error {
	return nil
}
