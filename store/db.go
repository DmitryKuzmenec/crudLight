package store

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/sqlite3"
	"github.com/golang-migrate/migrate/source/file"
	_ "github.com/mattes/migrate/source/file"
	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	*sql.DB
}

func New(dbPath string) (*DB, error) {
	sqliteDb, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("open sqlite DB failed: %w", err)
	}

	return &DB{sqliteDb}, nil
}

// Migrate launches migrations
func (db *DB) Migrate() error {
	driver, err := sqlite3.WithInstance(db.DB, &sqlite3.Config{})
	if err != nil {
		return fmt.Errorf("creating sqlite3 db driver failed: %w", err)
	}

	f, err := (&file.File{}).Open("file://../migrations")
	if err != nil {
		return fmt.Errorf("open migrations failed: %w", err)
	}

	m, err := migrate.NewWithInstance("file", f, "crud_light", driver)
	if err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}

	return m.Steps(1)
}
