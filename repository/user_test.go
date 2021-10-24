package repository

import (
	"testing"

	"github.com/DmitryKuzmenec/crudLight/models"
	"github.com/DmitryKuzmenec/crudLight/store"
	"github.com/stretchr/testify/assert"
)

var user *User

func init() {
	db, _ := store.New("./test_store.db")
	db.Migrate()
	user = NewUserAPI(db)
}

var in = &models.User{
	Name:      "Nemo",
	BirthDate: "2000-01-01",
}

func Test_Create(t *testing.T) {

	out, err := user.Create(in)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, int64(1), out.ID, "wrong user ID")

}
func Test_Get(t *testing.T) {
	out, err := user.Get(1)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, in.BirthDate, out.BirthDate, "wrong user birth date")
	assert.Equal(t, in.Name, out.Name, "wrong user Name")
	assert.Equal(t, int64(1), out.ID, "wrong user ID")
}

func Test_Update(t *testing.T) {
	err := user.Update(&models.User{ID: 1, Name: "Max"})
	if err != nil {
		t.Fatal(err)
	}
	out, err := user.Get(1)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, in.BirthDate, out.BirthDate, "wrong user birth date")
	assert.Equal(t, "Max", out.Name, "wrong user Name")
	assert.Equal(t, int64(1), out.ID, "wrong user ID")
}

func Test_Delete(t *testing.T) {
	err := user.Delete(1)
	if err != nil {
		t.Fatal(err)
	}

	user.db.Clean()
}
