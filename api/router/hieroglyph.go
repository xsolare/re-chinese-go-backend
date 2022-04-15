package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xsolare/re-chinese-go-backend/api/controllers"
)

type HieroglyphRouter struct{}

func (r *HieroglyphRouter) InitHieroglyphRouter(rg *gin.RouterGroup) {
	hieroglyphRouter := rg.Group("/hieroglyph")
	api := new(controllers.HieroglyphController)

	hieroglyphRouter.GET("/pinyin/:pinyin", api.GetByPinyin)
	hieroglyphRouter.GET("/name/:name", api.GetByName)
	hieroglyphRouter.GET("/:id", api.GetById)
	hieroglyphRouter.POST("/", api.NewHieroglyph)
	hieroglyphRouter.DELETE("/:id", api.DeleteHieroglyph)
}
