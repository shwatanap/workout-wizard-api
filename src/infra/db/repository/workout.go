package repository

import (
	"context"
	"database/sql"

	"github.com/shwatanap/workout-wizard-api/src/domain/entity"
	"github.com/shwatanap/workout-wizard-api/src/domain/repository"
	"github.com/shwatanap/workout-wizard-api/src/infra/db/sqlc"
)

type workoutRepository struct {
	queries *sqlc.Queries
}

func NewWorkoutRepository(db *sql.DB) repository.IWorkoutRepository {
	return &workoutRepository{queries: sqlc.New(db)}
}

func (wr *workoutRepository) Create(ctx context.Context, we *entity.Workout, menu_id int64) (int64, error) {
	err := wr.queries.CreateWorkout(ctx, sqlc.CreateWorkoutParams{
		MenuID: int32(menu_id),
		Name:   we.Name,
		Detail: we.Detail,
	})
	if err != nil {
		return 0, err
	}
	return wr.queries.GetCreatedWorkoutID(ctx)
}
