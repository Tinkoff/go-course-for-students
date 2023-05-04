package main

import (
	"context"

	pgxLogrus "github.com/jackc/pgx-logrus"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
	log "github.com/sirupsen/logrus"
)

func main() {
	logger := log.New()
	logger.SetLevel(log.InfoLevel)
	logger.SetFormatter(&log.TextFormatter{})

	config, err := pgxpool.ParseConfig("postgres://postgres:postgres@localhost:5433/pg_course")
	if err != nil {
		logger.WithError(err).Fatalf("can't parse pgxpool config")
	}

	config.ConnConfig.Tracer = &tracelog.TraceLog{
		Logger:   pgxLogrus.NewLogger(logger),
		LogLevel: tracelog.LogLevelDebug,
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		logger.WithError(err).Fatalf("can't create new pool")
	}
	defer pool.Close()

	user, err := GetUserByID(context.Background(), pool, 5165502)
	if err != nil {
		logger.WithError(err).Fatal("can't exec GetUserByID: %w", err)
	}

	logger.Infof("%+v", user)
}
