package main

import (
	"fmt"
	"os"

	"github.com/DmitryKuzmenec/crudLight/handlers"
	"github.com/DmitryKuzmenec/crudLight/repository"
	"github.com/DmitryKuzmenec/crudLight/restapi"
	"github.com/DmitryKuzmenec/crudLight/restapi/operations"
	"github.com/DmitryKuzmenec/crudLight/store"
	"github.com/go-openapi/loads"
	"github.com/jessevdk/go-flags"
	"go.uber.org/zap"
)

const DBPath = "./store.db"

func main() {

	log, _ := zap.NewProduction()
	defer log.Sync()

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Panic(err.Error())
	}

	api := operations.NewCrudLightAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	api.Logger = log.Sugar().Infof

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "API crudLight"
	parser.LongDescription = swaggerSpec.Spec().Info.Description
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Panic(err.Error())
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	// initialize database
	db, err := store.New(DBPath)
	if err != nil {
		log.Panic(fmt.Sprintf("database initializing failed: %s", err))
	}
	defer db.Close()

	// Do migrations
	if err := db.Migrate(); err != nil {
		log.Panic(fmt.Sprintf("database migration failed: %s", err))
	}

	// initialize user repo
	userRepo := repository.NewUserAPI(db)

	// initialize handlers
	api.UserGetHandler = handlers.UserGet(userRepo, log)
	api.UserCreateHandler = handlers.UserCreate(userRepo, log)
	api.UserUpdateHandler = handlers.UserUpdate(userRepo, log)
	api.UserDeleteHandler = handlers.UserDelete(userRepo, log)

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Panic(err.Error())
	}
}
