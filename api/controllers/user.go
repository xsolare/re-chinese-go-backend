package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xsolare/re-chinese-go-backend/api/models/dto"
	res "github.com/xsolare/re-chinese-go-backend/api/models/response"
	"github.com/xsolare/re-chinese-go-backend/global"
	utils "github.com/xsolare/re-chinese-go-backend/utils"
	"github.com/xsolare/re-chinese-go-backend/utils/jwtauth"
	"github.com/xsolare/re-chinese-go-backend/utils/jwtauth/jwtuser"
	"github.com/xsolare/re-chinese-go-backend/utils/response"
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

	if err := utils.Verify(req, utils.SignUpVerify); err != nil {
		global.GV_LOG.Error("Invalid data dto ", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err, user := userService.SignUp(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	j := jwtauth.NewJWT()
	claims := j.CreateClaims(jwtauth.BaseClaims{
		ID:       user.Id,
		Username: user.Username,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GV_LOG.Error("Fail token create!", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	response.OkWithData(res.UserJwt{
		User: user,
		Jwt:  token,
	}, c)
}

func (r *UserController) SignIn(c *gin.Context) {
	var req dto.SignIn
	_ = c.ShouldBindJSON(&req)

	if err := utils.Verify(req, utils.SignInVerify); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err, users := userService.SignIn(req)

	if err != nil {
		global.GV_LOG.Error("Incorrect password!", err)
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(users, c)
}

func (r *UserController) Auth(c *gin.Context) {
	//TODO Auth with token
	// var req dto.SignIn
	// _ = c.ShouldBindJSON(&req)

	// if err := utils.Verify(r, utils.SignInVerify); err != nil {
	// 	c.JSON(http.StatusBadRequest, err)
	// 	return
	// }

	// err, users := userService.SignIn(req)

	// if err != nil {
	// 	global.GV_LOG.Error("Nope!")
	// } else {
	// 	c.JSON(http.StatusOK, users)
	// }

	c.JSON(http.StatusOK, jwtuser.GetUserID(c))
}
