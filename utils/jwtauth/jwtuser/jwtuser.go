package jwtuser

import (
	"github.com/gin-gonic/gin"

	"github.com/xsolare/re-chinese-go-backend/global"
	jwtauth "github.com/xsolare/re-chinese-go-backend/utils/jwtauth"
)

func GetClaims(c *gin.Context) (*jwtauth.CustomClaims, error) {
	token := c.Request.Header.Get("x-token")
	j := jwtauth.NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		global.GV_LOG.Error("Getting parsing information from jwt from the Gin context failed, please check if there is an x token in the request header and if the statements match the specified structure.")
	}
	return claims, err
}

func GetUserId(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.ID
		}
	} else {
		waitUse := claims.(*jwtauth.CustomClaims)
		return waitUse.ID
	}
}

func GetUserInfo(c *gin.Context) *jwtauth.CustomClaims {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return nil
		} else {
			return cl
		}
	} else {
		waitUse := claims.(*jwtauth.CustomClaims)
		return waitUse
	}
}
