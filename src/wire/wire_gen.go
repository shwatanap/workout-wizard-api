// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"database/sql"
	"github.com/shwatanap/go-backend-template/src/infra/repository"
	"github.com/shwatanap/go-backend-template/src/presen/handler"
	"github.com/shwatanap/go-backend-template/src/usecase"
)

// Injectors from wire.go:

func InitTemplateHandler(driver *sql.DB) handler.TemplateHandler {
	iTemplateRepository := repository.NewTemplateRepository(driver)
	templateUsecase := usecase.NewTemplateUsecase(iTemplateRepository)
	templateHandler := handler.NewTemplateHandler(templateUsecase)
	return templateHandler
}
