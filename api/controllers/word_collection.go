package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/xsolare/re-chinese-go-backend/utils/response"
)

type WordCollectionController struct{}

/// 																	///

func (r *WordCollectionController) GetById(c *gin.Context) {
	id := c.Param("id")
	err, data := wordCollectionService.CollectionById(id)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(data, c)
}

///																											//
