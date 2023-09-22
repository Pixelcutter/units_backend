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
	// port := os.Getenv("PORT")
	fmt.Println("API is listening on port", app.Config.Port)

	srv := &http.Server {
		Addr: fmt.Sprintf(":%s", app.Config.Port),
		// TODO: add router
	}
	return srv.ListenAndServe()
}

func main() {
	err := godotenv.Load()
	if err != nil { 
		log.Fatal("Error loading .env")
	}

	cfg := Config {
		Port: os.Getenv("PORT"),
	}

	// TODO: connection to db

	app := &Application{
		Config: cfg,
	}

	err = app.Serve()
	if err != nil {
		log.Fatal(err)
	}
}