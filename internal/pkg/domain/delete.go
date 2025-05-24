package domain

import (
	"context"
	"fmt"
)

func (s *Service) Delete(id int) error {
	var ctx context.Context
	err := s.StorageRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return fmt.Errorf("element with ID %d not found", id)
}
