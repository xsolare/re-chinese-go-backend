package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xsolare/re-chinese-go-backend/configs"
	"github.com/xsolare/re-chinese-go-backend/global"
)

//? Cors
func Cors() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-Token")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}

//? CorsByRules
func CorsByRules() gin.HandlerFunc {
	if global.GV_CONFIG.Cors.Mode == "allow-all" {
		return Cors()
	}
	return func(c *gin.Context) {
		whitelist := checkCors(c.GetHeader("origin"))

		if whitelist != nil {
			c.Header("Access-Control-Allow-Origin", whitelist.AllowOrigin)
			c.Header("Access-Control-Allow-Headers", whitelist.AllowHeaders)
			c.Header("Access-Control-Allow-Methods", whitelist.AllowMethods)
			c.Header("Access-Control-Expose-Headers", whitelist.ExposeHeaders)
			if whitelist.AllowCredentials {
				c.Header("Access-Control-Allow-Credentials", "true")
			}
		}

		if whitelist == nil && global.GV_CONFIG.Cors.Mode == "strict-whitelist" && !(c.Request.Method == "GET" && c.Request.URL.Path == "/health") {
			c.AbortWithStatus(http.StatusForbidden)
		} else {
			if c.Request.Method == "OPTIONS" {
				c.AbortWithStatus(http.StatusNoContent)
			}
		}

		c.Next()
	}
}

func checkCors(currentOrigin string) *configs.CORSWhitelist {
	for _, whitelist := range global.GV_CONFIG.Cors.Whitelist {
		if currentOrigin == whitelist.AllowOrigin {
			return &whitelist
		}
	}
	return nil
}
