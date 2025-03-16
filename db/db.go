package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func DBConnect() (conn *pgx.Conn, ctx context.Context) {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	ctx = context.Background()
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	fmt.Println("Connected to database")
	return conn, ctx
}

func CreateTaskTable(conn *pgx.Conn, ctx context.Context) {
	_, err := conn.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS tasks (
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			description TEXT,
			status TEXT CHECK (status IN ('new', 'in_progress', 'done')) DEFAULT 'new',
			created_at TIMESTAMP DEFAULT now(),
			updated_at TIMESTAMP DEFAULT now()
		);
	`)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}
	log.Println("Table created or already exists.")
}
