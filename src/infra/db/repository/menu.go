package repository

import (
	"context"
	"database/sql"

	"github.com/shwatanap/workout-wizard-api/src/domain/entity"
	"github.com/shwatanap/workout-wizard-api/src/domain/repository"
	"github.com/shwatanap/workout-wizard-api/src/infra/db/sqlc"
)

type menuRepository struct {
	queries *sqlc.Queries
}

func NewMenuRepository(db *sql.DB) repository.IMenuRepository {
	return &menuRepository{queries: sqlc.New(db)}
}

func (mr *menuRepository) Create(ctx context.Context, me *entity.Menu) (err error) {
	err = mr.queries.CreateMenu(ctx, sqlc.CreateMenuParams{
		UserID:  me.UserID,
		Target:  me.Target,
		Comment: me.Comment,
	})
	return
}
