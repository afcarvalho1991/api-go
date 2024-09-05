package cpg

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	DB *pgxpool.Pool
}

var (
	pgInstance *Postgres
	pgOnce     sync.Once
)

func NewPG(ctx context.Context, connString string) (*Postgres, error) {
	pgOnce.Do(func() {
		DB, err := pgxpool.New(ctx, connString)
		if err != nil {
			// fmt.Errorf(fmt.Sprintf("unable to create connection pool: %w", err))
			return
		}
		pgInstance = &Postgres{DB}
	})

	return pgInstance, nil
}

func (pg *Postgres) Ping(ctx context.Context) error {
	return pg.DB.Ping(ctx)
}

func (pg *Postgres) Close() {
	pg.DB.Close()
}
