package handler

import (
	"github.com/gin-gonic/gin"
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
	_, err := mh.menuUsecase.Create(c)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
	})
}
