package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xsolare/re-chinese-go-backend/global"
)

type FinalController struct{}

/// 																	///

func (r *FinalController) GetFinals(c *gin.Context) {
	err, final := finalService.Finals()

	if err != nil {
		global.GV_LOG.Error("Nope!")
	}
	c.JSON(http.StatusOK, final)
}

// func GetFinals(c *gin.Context) {
// type Result struct {
// Id   int
// Name string
// }

// var result []Result

// global.GV_DB.Raw("SELECT id, name FROM finals").Scan(&result)
// logrus.Warn("result")

// var final models.Final
// err = global.GV_DB.First(&final).Error

// if err != nil {
// 	return err
// }

// c.JSON(http.StatusOK, result)

// if result != nil {
// 	c.JSON(http.StatusOK, result)
// } else {
// 	c.AbortWithStatus(http.StatusNotFound)
// }
// }
