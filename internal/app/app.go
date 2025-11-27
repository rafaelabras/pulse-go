package app

// The responsability of the app.go is to have a bunch of resources to use
// through the application, so it's viable to group structs here.

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rafaelabras/pulse-go/internal/api"
)

type Application struct {
	Logger      *log.Logger
	UserHandler *api.UserHandler
}

func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime)

	// stores

	// handlers
	userhandler := api.NewUserHandler()

	app := &Application{
		Logger:      logger,
		UserHandler: userhandler,
	}
	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "The service is healthy")
}
