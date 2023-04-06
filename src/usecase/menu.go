package usecase

import (
	"context"
	"log"

	"github.com/shwatanap/workout-wizard-api/src/domain/entity"
	"github.com/shwatanap/workout-wizard-api/src/domain/repository"
)

type MenuUsecase interface {
	Create(ctx context.Context) (*entity.Menu, error)
}

type menuUsecase struct {
	menuRepository         repository.IMenuRepository
	menuExternalRepository repository.IMenuExternalRepository
}

func NewMenuUsecase(mr repository.IMenuRepository, mer repository.IMenuExternalRepository) MenuUsecase {
	mu := menuUsecase{menuRepository: mr, menuExternalRepository: mer}
	return &mu
}

func (mu *menuUsecase) Create(ctx context.Context) (menu *entity.Menu, err error) {
	_, err = mu.menuExternalRepository.Create(ctx, 0, "モモ", "HelloWorld")
	// menu, err = mu.menuRepository.Create(ctx, 0, "モモ", "HelloWorld")
	log.Println("Hello")
	return
}
