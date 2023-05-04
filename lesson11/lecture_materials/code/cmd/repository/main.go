package main

import (
	"context"
	"pg-course/internal/repository"
	"pg-course/pkg/postgres"

	log "github.com/sirupsen/logrus"
)

func main() {
	logger := log.New()
	logger.SetLevel(log.InfoLevel)
	logger.SetFormatter(&log.TextFormatter{})

	pgConfig := postgres.Config{
		Host:     "localhost",
		Port:     5433,
		Database: "pg_course",
		User:     "postgres",
		Password: "postgres",
		MaxConns: 3,
		MinConns: 1,
	}

	postgresPool, err := postgres.NewPool(pgConfig, logger)
	if err != nil {
		logger.WithError(err).Fatal("can't create postgres pool")
	}

	repo := repository.NewRepository(postgresPool, logger)

	user, err := repo.FindUserByID(context.Background(), 555)
	if err != nil {
		logger.WithError(err).Fatal("can't exec FindUserByID: %w", err)
	}

	logger.Infof("%+v", user)
}
