package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xsolare/re-chinese-go-backend/api/controllers"
)

type PinyinRouter struct{}

/// 																	///

func (r *PinyinRouter) InitPinyinRouter(rg *gin.RouterGroup) {
	wordRouter := rg.Group("/pinyin")
	api := new(controllers.PinyinController)

	wordRouter.GET("/", api.GetPinyin)
	wordRouter.GET("/finals", api.GetFinals)
	wordRouter.GET("/initials", api.GetInitials)
	wordRouter.GET("/by-tone/:tone", api.GetByTone)
	wordRouter.GET("/by-inital-final", api.GetByInitalAndFinal)
}
