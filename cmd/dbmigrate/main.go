package main

import (
	"log"

	"github.com/cs-bird/cmd/dbmigrate/cfg"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	c := cfg.Get()

	if c.Migrate != "down" && c.Migrate != "up" {
		log.Println("-migrate accepts [up, down] values only")
		return
	}

	m, err := migrate.New("file://migrations", c.PsqlCfg.GetDBConnStr())
	if err != nil {
		log.Printf("%s", err)
		return
	}

	if c.Migrate == "up" {
		if err := m.Up(); err != nil {
			log.Printf("failed migrate up: %s\n", err)
			return
		}
	}

	if c.Migrate == "down" {
		if err := m.Down(); err != nil {
			log.Printf("failed migrate down: %s\n", err)
			return
		}
	}
}
