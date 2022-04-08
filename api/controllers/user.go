package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xsolare/re-chinese-go-backend/global"
)

type UserController struct{}

/// 																	///

func (r *UserController) GetAll(c *gin.Context) {
	err, users := userService.Users()

	if err != nil {
		global.GV_LOG.Error("Nope!")
	} else {
		c.JSON(http.StatusOK, users)
	}
}
