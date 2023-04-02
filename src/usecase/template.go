package usecase

import (
	"github.com/shwatanap/go-backend-template/src/domain/repository"
)

type TemplateUsecase interface{}

type templateUsecase struct {
	templateRepository repository.ITemplateRepository
}

func NewTemplateUsecase(tr repository.ITemplateRepository) TemplateUsecase {
	tu := templateUsecase{templateRepository: tr}
	return &tu
}
