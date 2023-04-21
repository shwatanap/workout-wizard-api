package repository

import (
	"context"

	"github.com/shwatanap/workout-wizard-api/src/domain/entity"
)

type IMenuRepository interface {
	Create(ctx context.Context, me *entity.Menu) (int64, error)
}

type IMenuExternalRepository interface {
	Create(ctx context.Context, user_id int, target, comment string) (*entity.Menu, error)
}
