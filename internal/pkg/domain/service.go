package domain

import (
	"canchitas-libres-field/internal/configuration"
	"context"
)

type StorageRepository interface {
	GetAll() ([]Field, error)
	GetByID(id int) (Field, error)
	Add(ctx context.Context, field Field) error
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, id int, fieldU Field) error
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
