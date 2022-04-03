package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xsolare/re-chinese-go-backend/api/models"
)

var users = []models.User{
	{Id: 1, Username: "injurka", Password: "injurka"},
	{Id: 2, Username: "vaetimidis", Password: "vaetimidis"},
}

type UserController struct{}

/// 																	///

func (r *UserController) GetUsers(c *gin.Context) {
	if users != nil {
		c.JSON(http.StatusOK, users)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}
