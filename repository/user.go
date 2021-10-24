package repository

import (
	"errors"
	"fmt"
	"strings"

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
func (u *User) Create(user *models.User) (*models.User, error) {
	if user == nil {
		return nil, errors.New("user is nil")
	}

	res, err := u.db.Exec("insert into users (name, birth_date) values ($1, $2)", user.Name, user.BirthDate)
	if err != nil {
		return nil, err
	}

	userID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	user.ID = userID

	return user, nil
}

// Get returns a user from db
func (u *User) Get(userID int64) (*models.User, error) {
	var user models.User
	res := u.db.QueryRow("select * from users where id = $1", userID)

	err := res.Err()
	if err != nil {
		return nil, err
	}

	res.Scan(&user.ID, &user.Name, &user.BirthDate)

	return &user, nil
}

// Update updates user in db
func (u *User) Update(user *models.User) error {
	if user == nil {
		return errors.New("user is nil")
	}

	var f []string
	var v []interface{}
	var i int

	if user.Name != "" {
		i++
		f = append(f, fmt.Sprintf("name = $%d", i))
		v = append(v, user.Name)
	}
	if user.BirthDate != "" {
		i++
		f = append(f, fmt.Sprintf("birth_date = $%d", i))
		v = append(v, user.BirthDate)
	}

	if len(f) == 0 {
		return nil
	}
	v = append(v, user.ID)

	q := fmt.Sprintf("update users set %s where id = $%d", strings.Join(f, ","), i+1)
	_, err := u.db.Exec(q, v...)

	return err
}

// Delete deletes user in db
func (u *User) Delete(userID int64) error {
	_, err := u.db.Exec("delete from users where id = $1", userID)
	return err
}
