package mapper

import (
	"canchitas-libres-field/internal/pkg/domain"
	"canchitas-libres-field/internal/pkg/infrastructure/web/dto"
)

func ToDomainField(dtoField dto.FieldDto) (domain.Field , error){
	field, _ := domain.NewField(dtoField.Type,dtoField.Price,dtoField.Name, dtoField.Status)
	return *field, nil
}
