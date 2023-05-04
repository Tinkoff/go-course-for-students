package main

import (
	"context"

	"github.com/Pallinder/go-randomdata"
	"github.com/jackc/pgx/v5"
	"github.com/schollz/progressbar/v3"
	log "github.com/sirupsen/logrus"
)

const count = 1000000

func main() {
	logger := log.New()
	logger.SetLevel(log.InfoLevel)
	logger.SetFormatter(&log.TextFormatter{})

	conn, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@localhost:5433/pg_course")
	if err != nil {
		logger.WithError(err).Fatalf("can't connect to pg")
	}
	defer conn.Close(context.Background())

	bar := progressbar.Default(count)

	for i := 0; i < count; i++ {
		bar.Add(1)
		name := randomdata.FirstName(randomdata.Male)
		date := randomdata.FullDateInRange("1950-01-01", "2020-01-01")

		_, err := conn.Exec(context.Background(), "INSERT INTO users (name, birthday) VALUES ($1, $2)", name, date)
		if err != nil {
			logger.WithError(err).Fatal("can't insert user")
		}
	}
}
