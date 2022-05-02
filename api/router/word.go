package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xsolare/re-chinese-go-backend/api/controllers"
)

type WordRouter struct{}

/// 																	///

func (r *WordRouter) InitWordRouter(rg *gin.RouterGroup) {
	wordRouter := rg.Group("/word")
	api := new(controllers.WordController)

	wordRouter.GET("/by-hieroglyph", api.GetByHieroglyph)

	wordRouter.
		POST("/", api.Add)

	// wordRouter.
	// 	Use(middleware.Auth()).
	// 	Use(middleware.Role([]int{models.ADMIN})).
	// 	DELETE("/:id", api.DeleteWord)
}
