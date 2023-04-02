//go:build wireinject
// +build wireinject

package wire

import (
	"database/sql"

	"github.com/google/wire"

	"github.com/shwatanap/go-backend-template/src/infra/repository"
	"github.com/shwatanap/go-backend-template/src/presen/handler"
	"github.com/shwatanap/go-backend-template/src/usecase"
)

func InitTemplateHandler(driver *sql.DB) handler.TemplateHandler {
	wire.Build(
		repository.NewTemplateRepository,
		usecase.NewTemplateUsecase,
		handler.NewTemplateHandler,
	)
	return nil
}
