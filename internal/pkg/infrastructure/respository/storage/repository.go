package storage

import (
	"canchitas-libres-field/internal/configuration"
	"canchitas-libres-field/internal/pkg/domain"

	"github.com/jmoiron/sqlx"
)

type Postgres struct {
	Configuration *configuration.Configuration
	*sqlx.DB
}
type Slice struct {
	Configuration *configuration.Configuration
	SliceArr      []domain.Field
}

func NewPostgresStorage(config *configuration.Configuration, db *sqlx.DB) *Postgres {
	return &Postgres{
		Configuration: config,
		DB:            db,
	}
}

func NewSliceStorage(config *configuration.Configuration) *Slice {
	return &Slice{
		Configuration: config,
		SliceArr:      make([]domain.Field, 0),
	}
}
