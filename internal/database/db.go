package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func ConnectDB() *pgx.Conn {

	dbURL, exists := os.LookupEnv("DATABASE_URL")
	if !exists {
		var user, password, local, port, dbname = "user", "admin", "localhost", "5432", "postgres"
		dbURL = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", user, password, local, port, dbname)
	}
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

	query := `
	create table if not exists tasks (
			id serial primary key,
			title text not null,
			description text,
			completed boolean default false,
			created_at TIMESTAMPTZ default NOW()
	);`
	_, err = conn.Exec(context.Background(), query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create table: %v\n", err)
		os.Exit(1)
	}

	println("Database connection and table check successful!")
	return conn
}
