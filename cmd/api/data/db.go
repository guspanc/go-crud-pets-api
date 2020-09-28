package data

import (
	"context"
	"database/sql"
	"github.com/guspanc/go-crud-pets-api/cmd/api/models"

	_ "github.com/go-sql-driver/mysql" // mysql dialect
)

// Datastore interface
type Datastore interface {
	Close(ctx context.Context) error
	GetPets(ctx context.Context) ([]models.Pet, error)
	AddPet(ctx context.Context, pet models.Pet) (models.Pet, error)
}

// DB struct (inherits Datastore)
type DB struct {
	*sql.DB
}

// NewDB returns datastore implementation
func NewDB(connStr string) (Datastore, error) {
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

// Close database connections
func (db *DB) Close(ctx context.Context) error {
	return db.Close(ctx)
}
