// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/SethukumarJ/CashierX-Auth-Service/pkg/api"
	"github.com/SethukumarJ/CashierX-Auth-Service/pkg/api/handler"
	"github.com/SethukumarJ/CashierX-Auth-Service/pkg/config"
	"github.com/SethukumarJ/CashierX-Auth-Service/pkg/db"
	"github.com/SethukumarJ/CashierX-Auth-Service/pkg/repository"
	"github.com/SethukumarJ/CashierX-Auth-Service/pkg/usecase"
)

// Injectors from wire.go:

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	transactionRepository := repository.NewTransactionRepository(gormDB)
	transactionUsecase := usecase.NewTransactionUsecase(transactionRepository)
	transactionHandler := handler.NewTransactionHandler(transactionUsecase)
	serverHTTP := http.NewServerHTTP(transactionHandler)
	return serverHTTP, nil
}
