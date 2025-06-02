package storage

import (
	"canchitas-libres-field/internal/pkg/domain"
	"context"
	"fmt"
)

const (
	queryInsertField = `
	INSERT INTO fields (name, type, price, status)
	VALUES ($1, $2, $3, $4);`

	querySelectAllFields = `
		SELECT 
			id,
			name,
			type,
			price,
			status
		FROM 
			fields;`

	querySelectFieldByID = `
		SELECT 
			id,
			name,
			type,
			price,
			status
		FROM 
			fields
		WHERE id = $1;`
	querySelectFieldsByType = `
		SELECT 
			id,
			name,
			type,
			price,
			status
		FROM 
			fields
		WHERE type = $1;`

	queryUpdateFieldName   = `UPDATE fields SET name = $1 WHERE id = $2;`
	queryUpdateFieldType   = `UPDATE fields SET type = $1 WHERE id = $2;`
	queryUpdateFieldPrice  = `UPDATE fields SET price = $1 WHERE id = $2;`
	queryUpdateFieldStatus = `UPDATE fields SET status = $1 WHERE id = $2;`

	queryDeleteField = `DELETE FROM fields WHERE id = $1;`
)

func (p *Postgres) GetAll() ([]domain.Field, error) {
	var fields []domain.Field
	err := p.DB.Select(&fields, querySelectAllFields)
	if err != nil {
		return nil, err
	}

	return fields, nil
}

func (p *Postgres) Add(ctx context.Context, field domain.Field) error {
	fmt.Println("in infrastructure layer we have a field with name: ", field.Name)

	//field.ID = uuid.New().String()

	tx, err := p.Begin()
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, queryInsertField, field.Name, field.Type, field.Price, field.Status)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to insert field: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (p *Postgres) GetByID(id string) (domain.Field, error) {
	var field domain.Field

	err := p.Get(&field, querySelectFieldByID, id)

	if err != nil {
		return domain.Field{}, fmt.Errorf("failed to get user by ID %s: %w", id, err)
	}
	return field, nil
}

func (p *Postgres) Delete(ctx context.Context, id string) error {
	tx, err := p.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	_, err = tx.ExecContext(ctx, queryDeleteField, id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete field: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	return nil
}
func (p *Postgres) Update(ctx context.Context, id string, fieldU domain.Field) error {
	currentField, err := p.GetByID(id)
	if err != nil {
		return fmt.Errorf("field not found: %w", err)
	}

	tx, err := p.Begin()
	if err != nil {
		return err
	}
	if fieldU.Name != "" {
		_, err = tx.ExecContext(ctx, queryUpdateFieldName, fieldU.Name, id)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to update field name: %w", err)
		}
	}
	if fieldU.Type != "" {
		_, err = tx.ExecContext(ctx, queryUpdateFieldType, fieldU.Type, id)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to update field type: %w", err)
		}
	}
	if fieldU.Price != 0 {
		_, err = tx.ExecContext(ctx, queryUpdateFieldPrice, fieldU.Price, id)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to update field price: %w", err)
		}
	}
	if fieldU.Status != currentField.Status {
		_, err = tx.ExecContext(ctx, queryUpdateFieldStatus, fieldU.Name, id)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to update field name: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	return nil
}

func (p *Postgres) GetByType(fieldType string) ([]domain.Field, error) {
	var fields []domain.Field
	//err := p.Get(&fields, querySelectFieldsByType, fieldType)
	err := p.Select(&fields, querySelectFieldsByType, fieldType)
	if err != nil {
		return []domain.Field{}, fmt.Errorf("failed to get field by type %s: %w", fieldType, err)
	}
	return fields, nil
}
