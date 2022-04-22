package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/xsolare/re-chinese-go-backend/utils/response"
)

type PinyinController struct{}

/// 																	///

func (r *PinyinController) GetByInitalAndFinal(c *gin.Context) {
	inital := c.Query("inital")
	final := c.Query("final")

	err, data := pinyinService.PinyinsByInitalAndFinal(inital, final)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(data, c)
}

func (r *PinyinController) GetByTone(c *gin.Context) {
	tone := c.Param("tone")

	err, data := pinyinService.ByTone(tone)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(data, c)
}

func (r *PinyinController) GetPinyin(c *gin.Context) {
	err, data := pinyinService.Full()

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(data, c)
}

func (r *PinyinController) GetFinals(c *gin.Context) {
	err, data := pinyinService.Finals()

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(data, c)
}

func (r *PinyinController) GetInitials(c *gin.Context) {
	err, data := pinyinService.Initials()

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(data, c)

}
