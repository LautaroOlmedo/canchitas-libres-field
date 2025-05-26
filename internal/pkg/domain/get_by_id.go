package domain

func (s *Service) GetByID(id string) (Field, error) {
	return s.StorageRepository.GetByID(id)
}
