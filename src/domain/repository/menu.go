package repository

import (
	"context"

	"github.com/shwatanap/workout-wizard-api/src/domain/entity"
)

type IMenuRepository interface {
	Create(ctx context.Context, me *entity.Menu) error
}

type IMenuExternalRepository interface {
	Create(ctx context.Context, user_id int, target, comment string) (*entity.Menu, error)
}
