package app

// The responsability of the app.go is to have a bunch of resources to use
// through the application, so it's viable to group structs here.

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rafaelabras/pulse-go/internal/api"
	"github.com/rafaelabras/pulse-go/internal/store"
	"github.com/rafaelabras/pulse-go/migrations"
)

type Application struct {
	Logger      *log.Logger
	DB          *sql.DB
	UserHandler *api.UserHandler
}

func NewApplication() (*Application, error) {
	pgDB, err := store.Open()

	if err != nil {
		return nil, err
	}

	err = store.MigrateFS(pgDB, migrations.FS, ".")

	if err != nil {
		panic(err)
	}

	logger := log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime)

	// stores

	// handlers
	userhandler := api.NewUserHandler()

	app := &Application{
		Logger:      logger,
		DB:          pgDB,
		UserHandler: userhandler,
	}
	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "The service is healthy")
}
