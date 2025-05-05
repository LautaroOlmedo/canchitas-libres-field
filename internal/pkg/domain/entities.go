package domain

import (
	"errors"
	"strings"
)

type Field struct {
	ID    int     `json:"id" db:"field_id"`
	Type  string  `json:"type" db:"type"`
	Price float64 `json:"price" db:"price"`
	Name string   `json:"name" db:"name"`
}

var (
	ErrMissingFieldParameter = errors.New("missing required field parameter")
	ErrInvalidFieldType      = errors.New("invalid field type")
	ErrInvalidPrice          = errors.New("price must be positive")
)

func NewField(fieldType string, price float64, name string) (*Field, error) {
	if fieldType == "" {
		return nil, ErrMissingFieldParameter
	}

	fieldType = strings.ToLower(fieldType)
	if fieldType != "futbol 5" && fieldType != "futbol 7" && fieldType != "futbol 11" {
		return nil, ErrInvalidFieldType
	}

	if price <= 0 {
		return nil, ErrInvalidPrice
	}

	return &Field{
		Type:  fieldType,
		Price: price,
		Name: name,
	}, nil
}
