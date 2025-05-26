package dto

import (
	"errors"
	"reflect"
	"strings"
)

var (
	ErrMissingParameter    = errors.New("missing parameter")
	ErrInvalidTypeVariable = errors.New("invalid type of variable")
)

type FieldDto struct {
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Price float64 `json:"price"`
	Status bool   `json:"status"`
}



// ValidateFieldCreateDto valida los campos obligatorios del DTO
func ValidateFieldCreateDto(name string, fieldType string, price float64) error {
	if strings.TrimSpace(name) == "" || strings.TrimSpace(fieldType) == "" || price <= 0 {
		return ErrMissingParameter
	}
	if reflect.TypeOf(name) != reflect.TypeOf("") || reflect.TypeOf(price) != reflect.TypeOf(0) || reflect.TypeOf(fieldType) != reflect.TypeOf("") {
		return ErrInvalidTypeVariable
	}
	return nil
}

// ValidateInputId valida que el ID sea válido
func ValidateInputId(id string) error {
	if id == "" {
		return ErrMissingParameter
	}
	return nil
}

func ValidateFieldUpdateDto(name string, fieldType string, price float64, status bool) error {
	if strings.TrimSpace(name) == "" || strings.TrimSpace(fieldType) == "" || price <= 0 {
		return ErrMissingParameter
	}
	// status es booleano, no hace falta validarlo como campo vacío.
	if reflect.TypeOf(name).Kind() != reflect.String ||
		reflect.TypeOf(fieldType).Kind() != reflect.String ||
		reflect.TypeOf(price).Kind() != reflect.Float64 ||
		reflect.TypeOf(status).Kind() != reflect.Bool {
		return ErrInvalidTypeVariable
	}
	return nil
}