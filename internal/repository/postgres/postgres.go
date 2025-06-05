package postgres

import (
	"context"
	"embed"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"final/internal/config"
	"final/internal/utils/observability/log"
	"final/migrations/migrator"
	pgmigrations "final/migrations/postgres"
)

type Postgres struct {
	connPool *pgxpool.Pool
	logger   *log.Logger
	cfg      *config.Postgres
}

func NewPostgres(
	ctx context.Context,
	cfg *config.Postgres,
	logger *log.Logger,
) *Postgres {
	conn := fmt.Sprintf(
		"postgres://%s:%s@%s:%v/%s?sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DbName, cfg.SSLMode,
	)

	pool, err := pgxpool.New(ctx, conn)
	if err != nil {
		logger.Error("connection to postgres failed", err)
		return nil
	}
	logger.Info("PostgreSQL connection pool initialized", "component", "postgres")

	migrats := []embed.FS{pgmigrations.FS}
	for _, fs := range migrats {
		if err := migrator.DoMigrate(fs, conn); err != nil {
			logger.Error("migration failed", err)
			return nil
		}
	}
	logger.Info("migrations applied successfully", "component", "postgres")

	return &Postgres{
		connPool: pool,
		logger:   logger,
		cfg:      cfg,
	}
}

func (p *Postgres) Close() {
	p.connPool.Close()
}
