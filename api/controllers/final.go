package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/xsolare/re-chinese-go-backend/utils/response"
)

type FinalController struct{}

/// 																	///

func (r *FinalController) GetAll(c *gin.Context) {
	err, data := finalService.Finals()

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(data, c)
}
