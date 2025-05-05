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

type FieldCreateDto struct {
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Price float64 `json:"price"`
}

type FieldDtoResponse struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Price float64 `json:"price"`
}

// ValidateFieldCreateDto valida los campos obligatorios del DTO
func ValidateFieldCreateDto(name string, fieldType string, price float64) error {
	if strings.TrimSpace(name) == "" || strings.TrimSpace(fieldType) == "" || price <= 0 {
		return ErrMissingParameter
	}
	if reflect.TypeOf(name) != reflect.TypeOf("") || reflect.TypeOf(fieldType) != reflect.TypeOf("") {
		return ErrInvalidTypeVariable
	}
	return nil
}

// ValidateInputId valida que el ID sea válido
func ValidateInputId(id int) error {
	if id == 0 {
		return ErrMissingParameter
	}
	if reflect.TypeOf(id) != reflect.TypeOf(1) {
		return ErrInvalidTypeVariable
	}
	return nil
}