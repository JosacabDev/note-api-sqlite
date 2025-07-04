package main

import (
	"database/sql"
	"github/JosacabDev/api-sqlite/internal/server"
	"github/JosacabDev/api-sqlite/pkg/logger"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func initDB(filepath, schemaPath string) (*sql.DB, error) {
	var err error
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	schemaBytes, err := os.ReadFile(schemaPath)
	if err != nil {
		return nil, err
	}
	schema := string(schemaBytes)
	_, err = db.Exec(schema)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	var err error

	logger.Init()

	db, err := initDB("notes.db", "db/schema.sql")
	if err != nil {
		logger.Fatal.Println(err)
	}
	logger.Info.Println("Database initialized successfully")
	defer db.Close()

	srv := server.NewServer(":8080", db)
	logger.Info.Println("Starting server on :8080")
	err = srv.Start()
	if err != nil {
		logger.Fatal.Println("Error starting server: ", err)
	}

}
