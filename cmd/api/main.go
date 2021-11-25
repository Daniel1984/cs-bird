package main

import (
	"log"

	"github.com/cs-bird/cmd/api/application"
	"github.com/cs-bird/cmd/api/router"
	"github.com/cs-bird/internals/exithandler"
	"github.com/cs-bird/internals/server"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load env vars: %s", err)
	}

	app, err := application.Get()
	if err != nil {
		log.Fatalf("failed initiate application: %s", err)
	}

	srv := server.
		Get().
		WithAddr(app.ApiPort).
		WithRouter(router.Get(app))

	go func() {
		log.Printf("starting server at %s", app.ApiPort)
		if err := srv.Start(); err != nil {
			log.Fatal(err.Error())
		}
	}()

	exithandler.Init(func() {
		if err := srv.Close(); err != nil {
			log.Println(err.Error())
		}

		if err := app.DB.Close(); err != nil {
			log.Println(err.Error())
		}
	})
}
