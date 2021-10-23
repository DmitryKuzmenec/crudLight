package repository

import (
	"github.com/DmitryKuzmenec/crudLight/models"
	"github.com/DmitryKuzmenec/crudLight/store"
)

type User struct {
	db *store.DB
}

func NewUserAPI(db *store.DB) *User {
	return &User{db: db}
}

// Create creates new user in db
func (u *User) Create(*models.User) (*models.User, error) {
	return &models.User{ID: 1, BirthDate: "123", Name: "Dmitry"}, nil
}

// Get returns a user from db
func (u *User) Get(userID int) (*models.User, error) {
	return nil, nil
}

// Update updates user in db
func (u *User) Update(*models.User) error {
	return nil
}

// Delete deletes user in db
func (u *User) Delete(userID int) error {
	return nil
}
