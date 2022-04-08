package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xsolare/re-chinese-go-backend/global"
)

type WordCollectionController struct{}

/// 																	///

func (r *WordCollectionController) GetById(c *gin.Context) {
	id := c.Param("id")
	err, data := wordCollectionService.CollectionById(id)

	if err != nil {
		global.GV_LOG.Error("Nope!")
	} else {
		c.JSON(http.StatusOK, data)
	}
}

///																											//
