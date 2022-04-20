package middleware

import (
	"time"

	handler "github.com/xsolare/re-chinese-go-backend/api/middleware/handler"
	g "github.com/xsolare/re-chinese-go-backend/global"
	jwt "github.com/xsolare/re-chinese-go-backend/utils/jwtauth"
)

//? Auth
func Auth() (*jwt.GinJWTMiddleware, error) {
	timeout := time.Hour

	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte(g.GV_CONFIG.JwtSecret),
		Timeout:     timeout,
		MaxRefresh:  time.Hour,
		PayloadFunc: handler.PayloadFunc,
		// IdentityHandler: handler.IdentityHandler,
		// Authenticator:   handler.Authenticator,
		// Authorizator:    handler.Authorizator,
		// Unauthorized:    handler.Unauthorized,
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

}
