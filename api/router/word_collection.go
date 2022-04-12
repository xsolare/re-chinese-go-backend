package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xsolare/re-chinese-go-backend/api/controllers"
)

type WordCollectionRouter struct{}

func (r *WordCollectionRouter) InitWordCollectionRouter(rg *gin.RouterGroup) {
	wordCollectionController := rg.Group("/word-collection")
	api := new(controllers.WordCollectionController)

	wordCollectionController.GET("/:id", api.GetById)
}
