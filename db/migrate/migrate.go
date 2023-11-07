package migrate

import (
	"database/sql"

	"github.com/ahmadfaizk/goose-app/config"
	"github.com/ahmadfaizk/goose-app/db"
	"github.com/ahmadfaizk/goose-app/internal/database"
	"github.com/pressly/goose/v3"
)

type Migrate struct {
	db  *sql.DB
	dir string
}

// New creates a new Migrate instance.
func New(cfg config.DatabaseConfig) (*Migrate, error) {
	dbConn, err := database.New(cfg)
	if err != nil {
		return nil, err
	}
	goose.SetBaseFS(db.MigrationFS)
	if err := goose.SetDialect("postgres"); err != nil {
		return nil, err
	}
	return &Migrate{db: dbConn.DB, dir: "migrations"}, nil
}

// Up runs all available migrations.
func (m *Migrate) Up() error {
	return goose.Up(m.db, m.dir)
}

// Down rolls back a single migration.
func (m *Migrate) Down() error {
	return goose.Down(m.db, m.dir)
}

// Reset rolls back all migrations.
func (m *Migrate) Reset() error {
	return goose.Reset(m.db, m.dir)
}

// Redo rolls back a single migration and then runs it again.
func (m *Migrate) Redo() error {
	return goose.Redo(m.db, m.dir)
}

// Version prints the current version of the database.
func (m *Migrate) Version() (int64, error) {
	return goose.GetDBVersion(m.db)
}
