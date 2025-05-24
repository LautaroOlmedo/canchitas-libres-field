package domain

import (
	"context"
	"errors"
	"strings"
)

var (
	ErrMissingFieldParameter = errors.New("missing required field parameter")
	ErrInvalidFieldType      = errors.New("invalid field type")
	ErrInvalidPrice          = errors.New("price must be positive")
	ErrInvalidStatus         = errors.New("invalid status value")
)

func (s *Service) Add(field Field) error {
	r := strings.ToLower(field.Type)
	if r == "" {
		return ErrMissingFieldParameter
	}

	
	if r != "futbol 5" && r != "futbol 7" && r != "futbol 11" {
		return ErrInvalidFieldType
	}
	if len(field.Name) > 30 {
		return errors.New("field name is too long")
	}

	if field.Price <= 0 {
		return ErrInvalidPrice
	}

	return s.StorageRepository.Add(context.Background(), field)

}
