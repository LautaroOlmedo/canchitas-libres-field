package domain

type Field struct {
	ID     string  `json:"id" db:"id"`
	Type   string  `json:"type" db:"type"`
	Price  float64 `json:"price" db:"price"`
	Name   string  `json:"name" db:"name"`
	Status bool    `json:"status" db:"status"`
}

func NewField(fieldType string, price float64, name string, status bool) (*Field, error) {

	return &Field{
		Type:   fieldType,
		Price:  price,
		Name:   name,
		Status: status,
	}, nil
}
