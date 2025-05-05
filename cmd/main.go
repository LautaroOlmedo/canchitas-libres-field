package main

import (
	"canchitas-libres-field/internal/configuration"
	database2 "canchitas-libres-field/internal/database"
	domain "canchitas-libres-field/internal/pkg/domain"
	"canchitas-libres-field/internal/pkg/infrastructure/respository/storage"
	"canchitas-libres-field/internal/pkg/infrastructure/web"
	"context"
	"fmt"
)

func main() {
	config, err := configuration.Load("../.env")
	if err != nil {
		panic(err)
	}

	// database connection
	db, err := database2.NewDBConnection(context.Background(), config)
	if err != nil {
		panic(err)
	}

	// repository layer
	postgresStorage := storage.NewPostgresStorage(config, db)

	// application layer (services layer)
	service := domain.NewService(config, postgresStorage)

	// infrastructure layer
	handler := web.NewHandler(service)
	server, err := web.NewServer(config, handler)
	if err != nil {
		fmt.Printf("error starting the server: %s\n", err)
		panic(err)
	}

	// Start application
	server.Start()
}
