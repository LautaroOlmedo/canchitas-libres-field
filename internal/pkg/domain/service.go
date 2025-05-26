package domain

import (
	"canchitas-libres-field/internal/configuration"
	"context"
)

type StorageRepository interface {
	GetAll() ([]Field, error)
	GetByID(id string) (Field, error)
	Add(ctx context.Context, field Field) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string, fieldU Field) error
	GetByType(fieldType string) (Field, error)
}

type Service struct {
	Config            *configuration.Configuration
	StorageRepository StorageRepository
}

func NewService(config *configuration.Configuration, storageRepository StorageRepository) *Service {
	return &Service{
		Config:            config,
		StorageRepository: storageRepository,
	}
}
