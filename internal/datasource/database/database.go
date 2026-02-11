package database

import (
	"context"
	"fmt"
	"simple_message/internal/config"
	"simple_message/internal/datasource/database/postgresql"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type (
	database struct {
		postgres *pgxpool.Pool
		redis    *redis.Client
	}

	Database interface {
		GetPostgress() *pgxpool.Pool
		Close(ctx context.Context)
	}
)

func New(ctx context.Context, cfg config.Config) (Database, error) {
	pg, err := postgresql.ConnectPostgres(ctx, cfg)
	if err != nil {
		return nil, err
	}

	re, err := NewRedisClientWithPing(ctx, cfg)
	if err != nil {
		pg.Close()
		return nil, err
	}

	return &database{
		postgres: pg,
		redis:    re,
	}, nil
}

func (db *database) GetPostgress() *pgxpool.Pool {
	return db.postgres
}

func (db *database) Close(ctx context.Context) {
	db.GetPostgress().Close()
}

func NewRedisClient(cfg config.Config) (*redis.Client, error) {
	// Создаем клиент Redis
	addr := fmt.Sprintf("%s:%s", cfg.GetRedisHost(), cfg.GetRedisPort())
	db, _ := strconv.Atoi(cfg.GetRedisDB())

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.GetRedisPassword(),
		DB:       db,
	})

	return client, nil
}

func NewRedisClientWithPing(ctx context.Context, cfg config.Config) (*redis.Client, error) {
	client, err := NewRedisClient(cfg)
	if err != nil {
		return nil, err
	}

	// Проверяем подключение
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("redis connection failed: %v", err)
	}

	return client, nil
}
