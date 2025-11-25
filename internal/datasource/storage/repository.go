package storage

import (
	"simple_messenger/internal/datasource/database"
	"simple_messenger/internal/datasource/storage/mongo_repo"
	"simple_messenger/internal/datasource/storage/postgres_repo"
)

type (
	repo struct {
		mongo mongo_repo.Storage
		pg    postgres_repo.Storage
	}

	Repo interface {
		Mongo() mongo_repo.Storage
		Postgres() postgres_repo.Storage
	}
)

func New(db database.Database) Repo {
	return &repo{
		mongo: mongo_repo.New(db.GetMongo()),
		pg:    postgres_repo.New(db.GetPostgress()),
	}
}

func (r *repo) Mongo() mongo_repo.Storage {
	return r.mongo
}

func (r *repo) Postgres() postgres_repo.Storage {
	return r.pg
}
