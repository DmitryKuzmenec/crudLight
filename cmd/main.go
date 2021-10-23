package main

import (
	"log"
	"os"

	"github.com/DmitryKuzmenec/crudLight/handlers"
	"github.com/DmitryKuzmenec/crudLight/repository"
	"github.com/DmitryKuzmenec/crudLight/restapi"
	"github.com/DmitryKuzmenec/crudLight/restapi/operations"
	"github.com/DmitryKuzmenec/crudLight/store"
	"github.com/go-openapi/loads"
	"github.com/jessevdk/go-flags"
)

const DBPath = "./store.db"

func main() {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewCrudLightAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "API crudLight"
	parser.LongDescription = swaggerSpec.Spec().Info.Description
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
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

	// Initial database
	db, err := store.New(DBPath)
	if err != nil {
		log.Fatal(err)
	}

	// Do migrations
	if err := db.Migrate(); err != nil {
		log.Fatal(err)
	}

	// initial user repo
	userRepo := repository.NewUserAPI(db)

	// initial handlers
	api.UserGetHandler = handlers.UserGet(userRepo)
	api.UserCreateHandler = handlers.UserCreate(userRepo)
	api.UserUpdateHandler = handlers.UserUpdate(userRepo)
	api.UserDeleteHandler = handlers.UserDelete(userRepo)

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
