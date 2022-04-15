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

func (r *PinyinController) GetByTone(c *gin.Context) {
	tone := c.Param("tone")

	err, pinyins := pinyinService.ByTone(tone)

	if err != nil {
		global.GV_LOG.Error("Nope!")
	} else {
		c.JSON(http.StatusOK, pinyins)
	}
}

func (r *PinyinController) GetPinyin(c *gin.Context) {
	err, pinyins := pinyinService.Full()

	if err != nil {
		global.GV_LOG.Error("Nope!")
	} else {
		c.JSON(http.StatusOK, pinyins)
	}
}

func (r *PinyinController) GetFinals(c *gin.Context) {
	err, finals := pinyinService.Finals()

	if err != nil {
		global.GV_LOG.Error("Nope!")
	} else {
		c.JSON(http.StatusOK, finals)
	}
}

func (r *PinyinController) GetInitials(c *gin.Context) {
	err, initials := pinyinService.Initials()

	if err != nil {
		global.GV_LOG.Error("Nope!")
	} else {
		c.JSON(http.StatusOK, initials)
	}
}
