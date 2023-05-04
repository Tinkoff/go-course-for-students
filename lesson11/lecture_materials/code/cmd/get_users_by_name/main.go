package main

import (
	"context"

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

	_, err = GetUsersByName(context.Background(), conn, "Anthony")
	if err != nil {
		logger.WithError(err).Fatal("can't exec GetUsersByName: %w", err)
	}

	//for _, user := range users {
	//	logger.Infof("%+v", user)
	//}
}
