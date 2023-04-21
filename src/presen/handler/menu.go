package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/shwatanap/workout-wizard-api/src/presen/request"
	"github.com/shwatanap/workout-wizard-api/src/presen/response"
	"github.com/shwatanap/workout-wizard-api/src/usecase"
)

type MenuHandler interface {
	Create(c *gin.Context)
}

type menuHandler struct {
	menuUsecase usecase.MenuUsecase
}

func NewMenuHandler(mu usecase.MenuUsecase) MenuHandler {
	return &menuHandler{menuUsecase: mu}
}

func (mh *menuHandler) Create(c *gin.Context) {
	var req request.CreateMenuRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	me, err := mh.menuUsecase.Create(c)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	cwrs := []response.CreateWorkoutResponse{}
	for _, we := range me.Workouts {
		cwr := response.CreateWorkoutResponse{
			ID:     we.ID,
			Name:   we.Name,
			Detail: we.Detail,
		}
		cwrs = append(cwrs, cwr)
	}

	res := response.CreateMenuResponse{
		ID:       me.ID,
		UserID:   me.UserID,
		Target:   me.Target,
		Comment:  me.Comment,
		Workouts: cwrs,
	}

	c.JSON(200, gin.H{
		"menus": res,
	})
}
