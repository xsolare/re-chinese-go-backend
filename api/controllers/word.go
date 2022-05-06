package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/xsolare/re-chinese-go-backend/api/models/dto"
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

func (r *WordController) GetById(c *gin.Context) {
	id := c.Param("id")
	err, data := wordService.ById(id)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(data, c)
}

func (r *WordController) Add(c *gin.Context) {
	var req dto.Word
	c.ShouldBindJSON(&req)

	err, data := wordService.AddWord(&req)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(data, c)
}
func (r *WordController) GetTranslateById(c *gin.Context) {
	id := c.Param("id")
	lang := c.Query("lang")

	err, data := wordService.TranslateById(id, lang)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(data, c)
}
