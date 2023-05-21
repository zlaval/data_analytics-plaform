package postgresdb

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type PostgresConfig struct {
	DSN string
}

type PostgresDB struct {
	DB *pgxpool.Pool
}

func (c *PostgresConfig) Connect() *PostgresDB {

	pool, err := pgxpool.New(context.TODO(), c.DSN)
	if err != nil {
		log.Panic("Cannot connect to postgres")
	}

	return &PostgresDB{
		DB: pool,
	}
}

func (p *PostgresDB) Close() {
	p.Close()
}
