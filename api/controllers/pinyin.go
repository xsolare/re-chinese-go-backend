package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xsolare/re-chinese-go-backend/global"
)

type PinyinController struct{}

/// 																	///

func (r *PinyinController) GetByInitalAndFinal(c *gin.Context) {
	inital := c.Query("inital")
	final := c.Query("final")

	err, words := pinyinService.PinyinsByInitalAndFinal(inital, final)

	if err != nil {
		global.GV_LOG.Error("Nope!")
	} else {
		c.JSON(http.StatusOK, words)
	}
}
