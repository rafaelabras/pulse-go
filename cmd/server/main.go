package main

import (
	"flag"
	"fmt"
	"github.com/rafaelabras/pulse-go/internal/app"
	"github.com/rafaelabras/pulse-go/internal/routes"
	"net/http"
	"time"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "go backend server port, with default value 8080")
	flag.Parse() // The flag allow us for dynamically replace values in cli,
	// in this case: go run main.go -port 8081 *example

	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	r := routes.SetupRoutes(app)
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	app.Logger.Printf("Starting golang backend server on port: %d", port)

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}

}
