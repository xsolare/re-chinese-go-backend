package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xsolare/re-chinese-go-backend/api/controllers"
)

type FinalRouter struct{}

/// 																	///

func (r *FinalRouter) InitFinalRouter(rg *gin.RouterGroup) {
	finalRouter := rg.Group("/final")
	api := new(controllers.FinalController)

	finalRouter.GET("/", api.GetAll)
}
