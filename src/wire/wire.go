//go:build wireinject
// +build wireinject

package wire

import (
	"database/sql"

	"github.com/google/wire"

	er "github.com/shwatanap/workout-wizard-api/src/infra/api/repository"
	"github.com/shwatanap/workout-wizard-api/src/infra/db/repository"
	"github.com/shwatanap/workout-wizard-api/src/presen/handler"
	"github.com/shwatanap/workout-wizard-api/src/usecase"
)

func InitMenuHandler(driver *sql.DB) handler.MenuHandler {
	wire.Build(
		repository.NewMenuRepository,
		repository.NewWorkoutRepository,
		er.NewMenuExternalRepository,
		usecase.NewMenuUsecase,
		handler.NewMenuHandler,
	)
	return nil
}
