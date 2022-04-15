package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xsolare/re-chinese-go-backend/global"
)

type HieroglyphCollectionController struct{}

/// 																	///

func (r *HieroglyphCollectionController) GetById(c *gin.Context) {
	err, hc := hieroglyphCollectionService.HieroglyphCollections()

	if err != nil {
		global.GV_LOG.Error("Nope!")
	} else {
		c.JSON(http.StatusOK, hc)
	}
}

func (r *HieroglyphCollectionController) AddHieroglyph(c *gin.Context) {
	err, hc := hieroglyphCollectionService.HieroglyphCollections()

	if err != nil {
		global.GV_LOG.Error("Nope!")
	} else {
		c.JSON(http.StatusOK, hc)
	}
}

func (r *HieroglyphCollectionController) DeleteHieroglyph(c *gin.Context) {
	err, hc := hieroglyphCollectionService.HieroglyphCollections()

	if err != nil {
		global.GV_LOG.Error("Nope!")
	} else {
		c.JSON(http.StatusOK, hc)
	}
}

func (r *HieroglyphCollectionController) CreateCollection(c *gin.Context) {
	err, hc := hieroglyphCollectionService.HieroglyphCollections()

	if err != nil {
		global.GV_LOG.Error("Nope!")
	} else {
		c.JSON(http.StatusOK, hc)
	}
}

func (r *HieroglyphCollectionController) DeleteCollection(c *gin.Context) {
	err, hc := hieroglyphCollectionService.HieroglyphCollections()

	if err != nil {
		global.GV_LOG.Error("Nope!")
	} else {
		c.JSON(http.StatusOK, hc)
	}
}

//
