package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/xsolare/re-chinese-go-backend/api/models/dto"
	res "github.com/xsolare/re-chinese-go-backend/api/models/response"
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
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(users, c)
}

func (r *UserController) SignUp(c *gin.Context) {
	var req dto.SignUp
	_ = c.ShouldBindJSON(&req)

	if err := utils.Verify(req, utils.SignUpVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err, user := userService.SignUp(req)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	j := jwtauth.NewJWT()
	claims := j.CreateClaims(jwtauth.BaseClaims{
		ID:       user.Id,
		Username: user.Username,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
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
	fmt.Printf("P - " + req.Password)
	fmt.Printf("U - " + req.Username)

	if err := utils.Verify(req, utils.SignInVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err, user := userService.SignIn(req)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	j := jwtauth.NewJWT()
	claims := j.CreateClaims(jwtauth.BaseClaims{
		ID:       user.Id,
		Username: user.Username,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(res.UserJwt{
		User: user,
		Jwt:  token,
	}, c)
}

func (r *UserController) Auth(c *gin.Context) {
	err, user := userService.Auth(jwtuser.GetUserId(c))

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	j := jwtauth.NewJWT()
	claims := j.CreateClaims(jwtauth.BaseClaims{
		ID:       user.Id,
		Username: user.Username,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(res.UserJwt{
		User: user,
		Jwt:  token,
	}, c)
}
