package main

import (
	"github.com/rafaelabras/pulse-go/internal/app"
)

func main() {
	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	app.Logger.Println("Starting server")
	
}
