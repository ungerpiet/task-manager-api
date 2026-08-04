package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func ConnectDB() *pgx.Conn {
	var user, password, local, port, dbname = "postgres", "admin", "localhost", "5432", "postgres"
	var dbURL string = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", user, password, local, port, dbname)

	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	err = conn.Ping(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to ping database: %v\n", err)
		os.Exit(1)
	}
	println("Database connection successful!")
	return conn
}
