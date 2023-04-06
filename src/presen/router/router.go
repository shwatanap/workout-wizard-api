package router

import (
	"github.com/gin-gonic/gin"

	"github.com/shwatanap/workout-wizard-api/src/presen/handler"
)

func InitRouting(menuHandler handler.MenuHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/menus/create", menuHandler.Create)
	return r
}
