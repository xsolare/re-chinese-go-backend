package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/xsolare/re-chinese-go-backend/utils/response"
)

type HieroglyphCollectionController struct{}

/// 																	///

func (r *HieroglyphCollectionController) GetById(c *gin.Context) {
	err, data := hieroglyphCollectionService.HieroglyphCollections()

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(data, c)
}

func (r *HieroglyphCollectionController) AddHieroglyph(c *gin.Context) {
	err, data := hieroglyphCollectionService.HieroglyphCollections()

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(data, c)
}

func (r *HieroglyphCollectionController) DeleteHieroglyph(c *gin.Context) {
	err, data := hieroglyphCollectionService.HieroglyphCollections()

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(data, c)
}

func (r *HieroglyphCollectionController) CreateCollection(c *gin.Context) {
	err, data := hieroglyphCollectionService.HieroglyphCollections()

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(data, c)
}

func (r *HieroglyphCollectionController) DeleteCollection(c *gin.Context) {
	err, data := hieroglyphCollectionService.HieroglyphCollections()

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(data, c)
}

//
