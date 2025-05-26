package domain

func (s *Service) GetByType(fieldType string) (Field, error) {
	return s.StorageRepository.GetByType(fieldType)
}