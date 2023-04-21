// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package sqlc

import (
	"context"
)

type Querier interface {
	CreateMenu(ctx context.Context, arg CreateMenuParams) error
	CreateWorkout(ctx context.Context, arg CreateWorkoutParams) error
	GetCreatedMenuID(ctx context.Context) (int64, error)
	GetCreatedWorkoutID(ctx context.Context) (int64, error)
	GetMenu(ctx context.Context, id int32) (interface{}, error)
}

var _ Querier = (*Queries)(nil)
