package app

// The responsability of the app.go is to have a bunch of resources to use
// through the application, so it's viable to group structs here.

import (
	"log"
	"os"
)

type Application struct {
	Logger *log.Logger
}

func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime)

	app := &Application{Logger: logger}
	return app, nil
}
