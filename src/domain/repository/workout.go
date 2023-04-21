package repository

import (
	"context"

	"github.com/shwatanap/workout-wizard-api/src/domain/entity"
)

type IWorkoutRepository interface {
	Create(ctx context.Context, we *entity.Workout, menu_id int64) (int64, error)
}
