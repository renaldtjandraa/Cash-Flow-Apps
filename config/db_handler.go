package config

import (
	utils "Cash-Flow-Apps/utils"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func Connect() (*pgx.Conn, error) {
	// Ambil DATABASE_URL dari environment variable
	connStr := utils.GetEenv("DATABASE_URL", "Error")
	if connStr == "Error" {
		log.Fatalf("DATABASE_URL environment variable not set.")
		return nil, fmt.Errorf("DATABASE_URL environment variable not set")
	}

	// Membuat koneksi ke database
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil, err
	}

	// Cek koneksi dengan melakukan query untuk melihat versi database
	var version string
	if err := conn.QueryRow(context.Background(), "SELECT version()").Scan(&version); err != nil {
		conn.Close(context.Background()) // Menutup koneksi
		log.Fatalf("Failed to select version: %v", err)
		return nil, err
	}

	log.Printf("Connected to database: %v", version)
	return conn, nil
}
