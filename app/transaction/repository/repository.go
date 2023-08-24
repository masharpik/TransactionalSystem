package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) (repo *Repository) {
	repo = &Repository{
		pool: pool,
	}
	return
}
