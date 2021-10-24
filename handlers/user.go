package handlers

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/DmitryKuzmenec/crudLight/models"
	"github.com/DmitryKuzmenec/crudLight/repository"
	api "github.com/DmitryKuzmenec/crudLight/restapi/operations/user"
	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"
)

var dateFormat = regexp.MustCompile(`^\d\d\d\d-\d\d-\d\d$`)

// UserCreate creates new user
func UserCreate(user *repository.User, log *zap.Logger) api.CreateHandlerFunc {
	return func(params api.CreateParams) middleware.Responder {
		u := params.Body
		if u == nil || u.Name == "" || u.BirthDate == "" || !dateFormat.MatchString(u.BirthDate) {
			log.Warn("wrong inbound data")
			return api.NewCreateBadRequest().WithPayload(
				&models.Error{
					Code:    http.StatusBadRequest,
					Message: "wrong inbound data",
				},
			)
		}

		newUser, err := user.Create(u)
		if err != nil {
			log.Warn(fmt.Sprintf("user creating failed: %s", err))
			return api.NewCreateInternalServerError().WithPayload(
				&models.Error{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			)
		}

		return api.NewCreateCreated().WithPayload(newUser)
	}
}

// UserGet returns user by id
func UserGet(user *repository.User, log *zap.Logger) api.GetHandlerFunc {
	return func(params api.GetParams) middleware.Responder {
		u, err := user.Get(params.ID)
		if err != nil {
			log.Warn(fmt.Sprintf("user getting failed: %s", err))
			return api.NewGetInternalServerError().WithPayload(
				&models.Error{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			)
		}

		return api.NewGetOK().WithPayload(u)
	}
}

// UserUpdate updates user by id
func UserUpdate(user *repository.User, log *zap.Logger) api.UpdateHandlerFunc {
	return func(params api.UpdateParams) middleware.Responder {
		u := params.Body
		u.ID = params.ID
		if err := user.Update(u); err != nil {
			log.Warn(fmt.Sprintf("user updating failed: %s", err))
			return api.NewUpdateInternalServerError().WithPayload(
				&models.Error{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			)
		}

		return api.NewUpdateOK()
	}
}

// UserDelete deletes user by id
func UserDelete(user *repository.User, log *zap.Logger) api.DeleteHandlerFunc {
	return func(params api.DeleteParams) middleware.Responder {
		if err := user.Delete(params.ID); err != nil {
			log.Warn(fmt.Sprintf("user deleting failed: %s", err))
			return api.NewDeleteInternalServerError().WithPayload(
				&models.Error{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			)
		}

		return api.NewDeleteOK()
	}
}
