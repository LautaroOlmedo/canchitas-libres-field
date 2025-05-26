package domain

import (
	"context"
	"errors"
	"strings"
)

func (s *Service) Update(id string, fieldU Field) error {
	var errIDNotFound = errors.New("ID not found")

	fieldU.Name = strings.TrimSpace(fieldU.Name)
	fieldU.Type = strings.TrimSpace(fieldU.Type)

	r := strings.ToLower(fieldU.Type)
	if r == "" {
		return ErrMissingFieldParameter
	}

	
	if r != "futbol 5" && r != "futbol 7" && r != "futbol 11" {
		return ErrInvalidFieldType
	}
	if len(fieldU.Name) > 30 {
		return errors.New("field name is too long")
	}

	if fieldU.Price <= 0 {
		return ErrInvalidPrice
	}

	fieldArray, errGetAll := s.StorageRepository.GetAll()
	if errGetAll != nil {
		return errGetAll
	}

	for i := range fieldArray {
		if fieldArray[i].ID == id {
			return s.StorageRepository.Update(context.Background(), id, fieldU)
		}
	}
	return errIDNotFound
}