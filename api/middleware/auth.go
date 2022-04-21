package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/xsolare/re-chinese-go-backend/utils/jwtauth"
	"github.com/xsolare/re-chinese-go-backend/utils/response"
)

//? JWTAuth Auth
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-token")
		j := jwtauth.NewJWT()
		claims, err := j.ParseToken(token)

		if err != nil {
			if err == jwtauth.TokenExpired {
				response.FailWithDetailed(gin.H{"reload": true}, "TokenExpired", c)
				c.Abort()
				return
			}
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
