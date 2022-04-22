package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/xsolare/re-chinese-go-backend/utils/response"
)

type WordController struct{}

/// 																	///

func (r *WordController) GetByHieroglyph(c *gin.Context) {
	hieroglyph := c.Query("hieroglyph")
	err, data := wordService.WordsByHieroglyph(hieroglyph)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(data, c)
}
