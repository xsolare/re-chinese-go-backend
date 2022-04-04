package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xsolare/re-chinese-go-backend/api/controllers"
)

type HieroglyphCollectionRouter struct{}

func (r *HieroglyphCollectionRouter) InitHieroglyphCollectionRouter(rg *gin.RouterGroup) {
	hieroglyphCollectionController := rg.Group("/hieroglyph_collection")
	api := new(controllers.HieroglyphCollectionController)

	hieroglyphCollectionController.GET("/", api.GetHieroglyphCollection)
}
