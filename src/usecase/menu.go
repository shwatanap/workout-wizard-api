package usecase

import (
	"context"

	"github.com/shwatanap/workout-wizard-api/src/domain/entity"
	"github.com/shwatanap/workout-wizard-api/src/domain/repository"
	"github.com/shwatanap/workout-wizard-api/src/utils"
)

type MenuUsecase interface {
	Create(ctx context.Context) (*entity.Menu, error)
}

type menuUsecase struct {
	workoutRepository      repository.IWorkoutRepository
	menuRepository         repository.IMenuRepository
	menuExternalRepository repository.IMenuExternalRepository
}

func NewMenuUsecase(wr repository.IWorkoutRepository, mr repository.IMenuRepository, mer repository.IMenuExternalRepository) MenuUsecase {
	mu := menuUsecase{
		workoutRepository:      wr,
		menuRepository:         mr,
		menuExternalRepository: mer,
	}
	return &mu
}

func (mu *menuUsecase) Create(ctx context.Context) (me *entity.Menu, err error) {
	// me, err = mu.menuExternalRepository.Create(ctx, 0, "モモ", "HelloWorld")
	// if err != nil {
	// 	return nil, err
	// }
	me = utils.ParseMenu("1. サイドプランク\n- 両手を肩幅に開いて、腕立て伏せのような体勢を作る。\n- 左側の腕を地面につけ、右側を上に伸ばす。\n- 右側の腕と足首を地面につけ、左腕を上に伸ばす。\n- この体勢で30秒〜1分間キープし、反対側も同じくやる。\n\n2. シットアップ\n- 背中をマットにつけて膝を立てる。\n- 両手を頭の後ろに組む。\n- 上半身を前に曲げ、腹筋を使って胸を膝に近づける。\n- 10〜15回を3セット程度行う。\n\n3. オブリークランチ\n- 背中をマットにつけて膝を立てる。\n- 両手を頭の後ろに組む。\n- 左足を右膝の上にのせる。\n- 上半身を前に曲げ、腕と逆の肘を左膝に近づける。\n- 反対側も同様にやる。\n- 10〜15回を3セット程度行う。\n\n4. ウッドチョップ\n- 両手を一直線に伸ばして、腰の高さに持ってくる。\n- 左足を踏み出しながら、腕を右側に回し上げる。\n- 反対側も同じく、気持ち良いくらいに回す。\n- 10〜15回を3セット程度行う。\n\n5. ナロースクワット\n- 足を肩幅よりも狭くして、膝を曲げてスクワットをする。\n- 膝が90度になるようにし、深くやると脇腹に効く。\n- 10〜15回を3セット程度行う。")
	menu_id, err := mu.menuRepository.Create(ctx, me)
	if err != nil {
		return nil, err
	}
	me.ID = int32(menu_id)

	for _, we := range me.Workouts {
		workout_id, err := mu.workoutRepository.Create(ctx, we, menu_id)
		if err != nil {
			return nil, err
		}
		we.ID = int32(workout_id)
	}
	return
}
