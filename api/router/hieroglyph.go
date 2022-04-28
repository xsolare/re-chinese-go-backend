package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xsolare/re-chinese-go-backend/api/controllers"
	"github.com/xsolare/re-chinese-go-backend/api/middleware"
	"github.com/xsolare/re-chinese-go-backend/api/models"
)

type HieroglyphRouter struct{}

func (r *HieroglyphRouter) InitHieroglyphRouter(rg *gin.RouterGroup) {
	hieroglyphRouter := rg.Group("/hieroglyph")
	api := new(controllers.HieroglyphController)

	hieroglyphRouter.GET("/pinyin/:pinyin", api.GetByPinyin)
	hieroglyphRouter.GET("/name/:name", api.GetByName)
	hieroglyphRouter.GET("/:id", api.GetById)
	hieroglyphRouter.
		Use(middleware.Auth()).
		Use(middleware.Role([]int{models.ADMIN})).
		POST("/", api.NewHieroglyph)

	hieroglyphRouter.
		Use(middleware.Auth()).
		Use(middleware.Role([]int{models.ADMIN})).
		DELETE("/:id", api.DeleteHieroglyph)
}
