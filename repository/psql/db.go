package psql

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type psqlDB struct {
	db *pgx.Conn
}

func New() *psqlDB {
	urlExample := "postgres://myuser:secret@localhost:5431/url_shortening_service"
	db, err := pgx.Connect(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	fmt.Print("CONNECCC")

	return &psqlDB{db: db}
}