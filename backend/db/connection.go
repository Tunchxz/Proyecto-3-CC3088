package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

func Connect() error {
	url := os.Getenv("DB_URL")
	var err error
	Conn, err = pgx.Connect(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error connecting to DB: %w", err)
	}
	fmt.Println("✅ Conectado a PostgreSQL")
	return nil
}
