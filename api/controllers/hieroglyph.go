package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xsolare/re-chinese-go-backend/api/models/dto"
	"github.com/xsolare/re-chinese-go-backend/global"
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
	err, hc := hieroglyphService.ById(id)

	if err != nil {
		global.GV_LOG.Error("Nope!")
	} else {
		c.JSON(http.StatusOK, hc)
	}
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
	err, hc := hieroglyphService.ByName(name)

	if err != nil {
		global.GV_LOG.Error("Nope!")
	} else {
		c.JSON(http.StatusOK, hc)
	}
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
	err, hc := hieroglyphService.ByPinyin(pinyin)

	if err != nil {
		global.GV_LOG.Error("Nope!")
	} else {
		c.JSON(http.StatusOK, hc)
	}
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
		global.GV_LOG.Error("Error create Hieroglyph!")
	} else {
		c.JSON(http.StatusOK, data)
	}
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
		global.GV_LOG.Error("Error delete Hieroglyph!")
	} else {
		c.JSON(http.StatusOK, data)
	}
}

///																											//
