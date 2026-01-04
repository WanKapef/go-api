package handler_test

import (
	"database/sql"
	"testing"

	"github.com/WanKapef/go-api/internal/database"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
)

func newTestDB(t *testing.T) *sql.DB {
	t.Helper()

	// banco em memória isolado
	dsn := "file::memory:?mode=memory&cache=private"

	db := database.ConnectSQLite(dsn)

	driver, err := sqlite.WithInstance(db, &sqlite.Config{})
	if err != nil {
		t.Fatalf("sqlite driver error: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://../../../migrations",
		"sqlite3", driver,
	)
	if err != nil {
		t.Fatalf("migrate init error: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		t.Fatalf("migrate up error: %v", err)
	}

	return db
}
