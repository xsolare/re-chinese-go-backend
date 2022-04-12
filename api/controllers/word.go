package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xsolare/re-chinese-go-backend/global"
)

type WordController struct{}

/// 																	///

func (r *WordController) GetByHieroglyph(c *gin.Context) {
	hieroglyph := c.Query("hieroglyph")
	err, words := wordService.WordsByHieroglyph(hieroglyph)

	if err != nil {
		global.GV_LOG.Error("Nope!")
	} else {
		c.JSON(http.StatusOK, words)
	}
}
