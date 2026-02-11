package postgresql

import (
	"context"
	"fmt"
	"simple_message/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectPostgres(ctx context.Context, cfg config.Config) (pool *pgxpool.Pool, err error) {
	dns := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.GetPGUser(),
		cfg.GetPGPassword(),
		cfg.GetPGHost(),
		cfg.GetPGPort(),
		cfg.GetPGName(),
	)
	pool, err = pgxpool.New(ctx, dns)
	if err != nil {
		return nil, fmt.Errorf("database connection failed: %v", err)
	}
	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("database ping failed: %v", err)
	}
	return pool, nil
}
