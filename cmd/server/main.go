package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

type Application struct {
	Config Config
}

func (app *Application) Serve() error {
	fmt.Println("API is listening on port", app.Config.Port)

	srv := &http.Server {
		Addr: fmt.Sprintf(":%s", app.Config.Port),
		// TODO: add router
	}
	
	// http.Server.ListenAndServe() returns an error
	return srv.ListenAndServe()
}

func main() {
	// loading .env vars
	if err := godotenv.Load(); err != nil { 
		log.Fatal("Error loading .env")
	}

	cfg := Config {
		Port: os.Getenv("PORT"),
	}

	// TODO: connection to db

	app := &Application{
		Config: cfg,
	}

	// starting server
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}