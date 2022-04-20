package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xsolare/re-chinese-go-backend/api/models/dto"
	"github.com/xsolare/re-chinese-go-backend/global"
	utils "github.com/xsolare/re-chinese-go-backend/utils"
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

func (r *UserController) SignUp(c *gin.Context) {
	var req dto.SignUp
	_ = c.ShouldBindJSON(&req)

	if err := utils.Verify(r, utils.SignUpVerify); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err, users := userService.SignUp(req)

	if err != nil {
		global.GV_LOG.Error("Nope!")
	} else {
		c.JSON(http.StatusOK, users)
	}
}
