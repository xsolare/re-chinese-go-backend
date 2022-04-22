package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/xsolare/re-chinese-go-backend/api/models/dto"
	"github.com/xsolare/re-chinese-go-backend/utils/response"
)

type HieroglyphController struct {
}

/// 																	///

/*
 * GetById
 * @Accept application/json
 * @Produce application/json
 * @Param data body {id} true "..."
 * @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "..."
 * @Router /hieroglyph/:id [get]
 */
func (r *HieroglyphController) GetById(c *gin.Context) {
	id := c.Param("id")
	err, data := hieroglyphService.ById(id)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(data, c)
}

/*
 * GetByName
 * @Accept application/json
 * @Produce application/json
 * @Param data body {name} true "..."
 * @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "..."
 * @Router /hieroglyph/name/:name [get]
 */
func (r *HieroglyphController) GetByName(c *gin.Context) {
	name := c.Param("name")
	err, data := hieroglyphService.ByName(name)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(data, c)
}

/*
 * GetByPinyin
 * @Accept application/json
 * @Produce application/json
 * @Param data body {pinyin} true "..."
 * @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "..."
 * @Router /hieroglyph/pinyin/:name [get]
 */
func (r *HieroglyphController) GetByPinyin(c *gin.Context) {
	pinyin := c.Param("pinyin")
	err, data := hieroglyphService.ByPinyin(pinyin)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(data, c)
}

/*
 * NewHieroglyph
 * @Accept application/json
 * @Produce application/json
 * @Param data body request.Hieroglyph true "..."
 * @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "..."
 * @Router /hieroglyph/ [post]
 */
func (r *HieroglyphController) NewHieroglyph(c *gin.Context) {
	var req dto.Hieroglyph
	c.ShouldBindJSON(&req)

	err, data := hieroglyphService.AddHieroglyph(&req)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(data, c)
}

/*
 * DeleteHieroglyph
 * @Accept application/json
 * @Produce application/json
 * @Param data body request.Hieroglyph true "..."
 * @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "..."
 * @Router /hieroglyph/:id [delete]
 */
func (r *HieroglyphController) DeleteHieroglyph(c *gin.Context) {
	id := c.Param("id")
	err, data := hieroglyphService.DeleteHieroglyph(id)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(data, c)
}

///																											//
