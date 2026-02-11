package storage

import (
	"simple_message/internal/datasource/database"
	"simple_message/internal/datasource/storage/postgres_repo"
)

type (
	repo struct {
		pg postgres_repo.Storage
	}

	Repo interface {
		Postgres() postgres_repo.Storage
	}
)

func New(db database.Database) Repo {
	return &repo{
		pg: postgres_repo.New(db.GetPostgress()),
	}
}

func (r *repo) Postgres() postgres_repo.Storage {
	return r.pg
}
