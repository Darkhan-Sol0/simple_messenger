package database

import (
	"context"
	"simple_messenger/internal/config"
	"simple_messenger/internal/datasource/database/mongodb"
	"simple_messenger/internal/datasource/database/postgresql"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type (
	database struct {
		postgres *pgxpool.Pool
		mongo    *mongo.Client
	}

	Database interface {
		GetPostgress() *pgxpool.Pool
		GetMongo() *mongo.Client
		Close(ctx context.Context)
	}
)

func New(ctx context.Context, cfg config.Config) (Database, error) {
	pg, err := postgresql.ConnectPostgres(ctx, cfg)
	if err != nil {
		return nil, err
	}
	mg, err := mongodb.ConnectMongo(ctx, cfg)
	if err != nil {
		return nil, err
	}
	return &database{
		postgres: pg,
		mongo:    mg,
	}, nil
}

func (db *database) GetPostgress() *pgxpool.Pool {
	return db.postgres
}

func (db *database) GetMongo() *mongo.Client {
	return db.mongo
}

func (db *database) Close(ctx context.Context) {
	db.GetPostgress().Close()
	db.GetMongo().Disconnect(ctx)
}
