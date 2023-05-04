package queries

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Queries struct {
	pool *pgxpool.Pool
}

func New(pgxPool *pgxpool.Pool) *Queries {
	return &Queries{pool: pgxPool}
}
