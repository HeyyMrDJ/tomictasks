package database

import (
	"database/sql"
	"log"
	"log/slog"

	_ "github.com/mattn/go-sqlite3"
)

func CreateTable(db *sql.DB) {
	sqlStmt := `
		CREATE TABLE IF NOT EXISTS lists (
		    id INTEGER PRIMARY KEY AUTOINCREMENT,
		    name TEXT NOT NULL,
		    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		slog.Debug("%q: %s\n", err, sqlStmt)
		return
	}

	sqlStmt = `
		CREATE TABLE IF NOT EXISTS tasks (
		    id INTEGER PRIMARY KEY AUTOINCREMENT,
		    list_id INTEGER,
		    title TEXT NOT NULL,
		    description TEXT,
		    due_date DATETIME,
		    completed BOOLEAN DEFAULT FALSE,
		    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		    FOREIGN KEY (list_id) REFERENCES lists(id) ON DELETE CASCADE
		);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		slog.Error("%q: %s\n", err, sqlStmt)
		return
	}
	CreateDefaultList(db)

}

func CreateDefaultList(db *sql.DB) {
	var exists bool
	sqlStmt := `
		SELECT EXISTS(SELECT 1 FROM lists WHERE name = ?)
	`
	err := db.QueryRow(sqlStmt, "default").Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		slog.Debug("Error checking for default list: %s\n", err)
		return
	}

	if !exists {
		sqlStmt = `
            INSERT INTO lists(name) VALUES(?);
        `
		_, err := db.Exec(sqlStmt, "default")
		if err != nil {
			log.Printf("Error creating default list: %s\n", err)
			return
		}
		slog.Debug("Default list created successfully")
	} else {
		slog.Debug("Default list already exists")
	}
}

func ConnectDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./tomictasks.db")
	if err != nil {
		slog.Error(err.Error())
	}

	return db
}
