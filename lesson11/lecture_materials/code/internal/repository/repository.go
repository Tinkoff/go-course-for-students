package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"pg-course/internal/domain"
	"pg-course/internal/repository/queries"
)

type repo struct {
	*queries.Queries
	pool   *pgxpool.Pool
	logger logrus.FieldLogger
}

func NewRepository(pgxPool *pgxpool.Pool, logger logrus.FieldLogger) Repository {
	return &repo{
		Queries: queries.New(pgxPool),
		pool:    pgxPool,
		logger:  logger,
	}
}

type Repository interface {
	FindUserByID(ctx context.Context, id int) (*domain.User, error)
}
