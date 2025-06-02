package domain

import (
	"context"
	"fmt"
)

func (s *Service) Delete(id string) error {
	fields, err := s.StorageRepository.GetAll()
	if err != nil {
		return err
	}

	for i := range fields {
		if fields[i].ID == id {
			return s.StorageRepository.Delete(context.Background(), id)
		}
	}

	return fmt.Errorf("element with ID %s not found", id)
}
