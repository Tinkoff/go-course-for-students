package main

import (
	"context"

	pgxLogrus "github.com/jackc/pgx-logrus"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/tracelog"
	log "github.com/sirupsen/logrus"
)

func main() {
	logger := log.New()
	logger.SetLevel(log.InfoLevel)
	logger.SetFormatter(&log.TextFormatter{})

	config, err := pgx.ParseConfig("postgres://postgres:postgres@localhost:5433/pg_course")
	if err != nil {
		logger.WithError(err).Fatalf("can't parse pgx config")
	}

	config.Tracer = &tracelog.TraceLog{
		Logger:   pgxLogrus.NewLogger(logger),
		LogLevel: tracelog.LogLevelDebug,
	}

	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		logger.WithError(err).Fatalf("can't connect to pg")
	}
	defer conn.Close(context.Background())

	user, err := GetUserByID(context.Background(), conn, 18210843)
	if err != nil {
		logger.WithError(err).Fatal("can't exec GetUserByID: %w", err)
	}

	logger.Infof("%+v", user)
}
