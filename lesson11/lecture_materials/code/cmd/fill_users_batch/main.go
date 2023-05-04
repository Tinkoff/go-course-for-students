package main

import (
	"context"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/jackc/pgx/v5"
	log "github.com/sirupsen/logrus"
)

func main() {
	logger := log.New()
	logger.SetLevel(log.InfoLevel)
	logger.SetFormatter(&log.TextFormatter{})

	conn, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@localhost:5433/pg_course")
	if err != nil {
		logger.WithError(err).Fatalf("can't connect to pg")
	}
	defer conn.Close(context.Background())

	start := time.Now()
	batch := &pgx.Batch{}
	for i := 0; i < 1000000; i++ {
		batch.Queue("INSERT INTO users (name, birthday) VALUES ($1, $2)",
			randomdata.FirstName(randomdata.Male), randomdata.FullDateInRange("1950-01-01", "2020-01-01"))
	}

	br := conn.SendBatch(context.Background(), batch)
	defer br.Close()

	if _, err := br.Exec(); err != nil {
		logger.WithError(err).Fatal("can't exec pg batch: %w", err)
	}

	logger.Infof("elapsed: %s", time.Now().Sub(start).String())
}
