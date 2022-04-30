package controllers

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/xsolare/re-chinese-go-backend/api/models"
	res "github.com/xsolare/re-chinese-go-backend/api/models/response"
	"github.com/xsolare/re-chinese-go-backend/global"
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

	s_pinyin, err := global.GV_REDIS.Get("pinyin")
	if err != nil {
		response.FailWithMessage(err.Error(), c)

		err, data := pinyinService.Full()
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		redis_data, _ := json.Marshal(data)
		global.GV_REDIS.Set("pinyin", string(redis_data), -1)
		response.OkWithData(data, c)
		return
	}

	var pinyin []res.Pinyin
	json.Unmarshal([]byte(s_pinyin), &pinyin)

	response.OkWithData(pinyin, c)

	// err, data := pinyinService.Full()
	// if err != nil {
	// 	response.FailWithMessage(err.Error(), c)
	// 	return
	// }
	// response.OkWithData(data, c)
}

func (r *PinyinController) GetFinals(c *gin.Context) {
	s_finals, err := global.GV_REDIS.Get("finals")
	if err != nil {
		response.FailWithMessage(err.Error(), c)

		err, data := pinyinService.Finals()
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		redis_data, _ := json.Marshal(data)
		global.GV_REDIS.Set("finals", string(redis_data), -1)
		response.OkWithData(data, c)
		return
	}

	var finals []models.Final
	json.Unmarshal([]byte(s_finals), &finals)

	response.OkWithData(finals, c)
}

func (r *PinyinController) GetInitials(c *gin.Context) {
	s_initials, err := global.GV_REDIS.Get("initials")
	if err != nil {
		response.FailWithMessage(err.Error(), c)

		err, data := pinyinService.Initials()
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		redis_data, _ := json.Marshal(data)
		global.GV_REDIS.Set("initials", string(redis_data), -1)
		response.OkWithData(data, c)
		return
	}

	var initials []models.Initial
	json.Unmarshal([]byte(s_initials), &initials)

	response.OkWithData(initials, c)
}
