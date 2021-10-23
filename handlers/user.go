package handlers

import (
	"github.com/DmitryKuzmenec/crudLight/repository"
	api "github.com/DmitryKuzmenec/crudLight/restapi/operations/user"
	"github.com/go-openapi/runtime/middleware"
)

// UserCreate creates new user
func UserCreate(user *repository.User) api.CreateHandlerFunc {
	return func(params api.CreateParams) middleware.Responder {
		u := params.Body
		if u == nil || u.Name == "" || u.BirthDate == "" {
			return api.NewCreateBadRequest()
		}
		newUser, err := user.Create(u)
		if err != nil {
			return api.NewCreateInternalServerError()
		}

		return api.NewCreateCreated().WithPayload(newUser)
	}
}

// UserGet returns user by id
func UserGet(user *repository.User) api.GetHandlerFunc {
	return func(params api.GetParams) middleware.Responder {
		u, err := user.Get(int(params.ID))
		if err != nil {
			return api.NewGetInternalServerError()
		}

		return api.NewGetOK().WithPayload(u)
	}
}

// UserUpdate updates user by id
func UserUpdate(user *repository.User) api.UpdateHandlerFunc {
	return func(params api.UpdateParams) middleware.Responder {
		u := params.Body
		u.ID = params.ID
		if err := user.Update(u); err != nil {
			return api.NewUpdateInternalServerError()
		}

		return api.NewUpdateOK()
	}
}

// UserDelete deletes user by id
func UserDelete(user *repository.User) api.DeleteHandlerFunc {
	return func(params api.DeleteParams) middleware.Responder {
		if err := user.Delete(int(params.ID)); err != nil {
			return api.NewDeleteInternalServerError()
		}

		return api.NewDeleteOK()
	}
}
