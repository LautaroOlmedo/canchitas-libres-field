package database

import (
	"canchitas-libres-field/internal/configuration"
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	notDriverProvides = errors.New("not driver provides")
)

func NewDBConnection(ctx context.Context, config *configuration.Configuration) (*sqlx.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DATABASE.DB_HOST,
		parseInt(config.DATABASE.DB_PORT),
		config.DATABASE.DB_USERNAME,
		config.DATABASE.DB_PASSWORD,
		config.DATABASE.DB_NAME)

	fmt.Println("conected to database")
	return sqlx.ConnectContext(ctx, "postgres", connectionString)
}

func parseInt(value string) int {
	val, err := strconv.Atoi(value)
	if err != nil {
		return 0.0
	}
	return val
}
