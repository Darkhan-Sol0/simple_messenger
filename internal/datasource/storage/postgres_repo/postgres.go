package postgres_repo

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type (
	pg_repo struct {
		client *pgxpool.Pool
	}

	Storage interface {
	}
)

var table = "chats"

func New(client *pgxpool.Pool) Storage {
	return &pg_repo{
		client: client,
	}
}
