package jwtauth

import (
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const JwtPayloadKey = "JWT_PAYLOAD"

type MapClaims map[string]interface{}

type GinJWTMiddleware struct {
	//? Realm name to display to the user. Required.
	Realm string

	//? signing algorithm - possible values are HS256, HS384, HS512
	//? Optional, default is HS256.
	SigningAlgorithm string

	//? Secret key used for signing. Required.
	Key []byte

	//? Duration that a jwt token is valid. Optional, defaults to one hour.
	Timeout time.Duration

	//? This field allows clients to refresh their token until MaxRefresh has passed.
	//? Note that clients can refresh their token in the last moment of MaxRefresh.
	//? This means that the maximum validity timespan for a token is TokenTime + MaxRefresh.
	//? Optional, defaults to 0 meaning not refreshable.
	MaxRefresh time.Duration

	//? Callback function that should perform the authentication of the user based on login info.
	//? Must return user data as user identifier, it will be stored in Claim Array. Required.
	//? Check error (e) to determine the appropriate error message.
	Authenticator func(c *gin.Context) (interface{}, error)

	//? Callback function that should perform the authorization of the authenticated user. Called
	//? only after an authentication success. Must return true on success, false on failure.
	//? Optional, default to success.
	Authorizator func(data interface{}, c *gin.Context) bool

	//? Callback function that will be called during login.
	//? Using this function it is possible to add additional payload data to the webtoken.
	//? The data is then made available during requests via c.Get("JWT_PAYLOAD").
	//? Note that the payload is not encrypted.
	//? The attributes mentioned on jwt.io can't be used as keys for the map.
	//? Optional, by default no additional data will be set.
	PayloadFunc func(data interface{}) MapClaims

	//? User can define own Unauthorized func.
	Unauthorized func(*gin.Context, int, string)

	//? User can define own LoginResponse func.
	LoginResponse func(*gin.Context, int, string, time.Time)

	//? User can define own RefreshResponse func.
	RefreshResponse func(*gin.Context, int, string, time.Time)

	//? Set the identity handler function
	IdentityHandler func(*gin.Context) interface{}

	//? Set the identity key
	IdentityKey string

	//? username
	NiceKey string

	DataScopeKey string

	//? rolekey
	RKey string

	//? roleId
	RoleIdKey string

	RoleKey string

	//? roleName
	RoleNameKey string

	//? TokenLookup is a string in the form of "<source>:<name>" that is used
	//? to extract token from the request.
	//? Optional. Default value "header:Authorization".
	//? Possible values:
	//? - "header:<name>"
	//? - "query:<name>"
	//? - "cookie:<name>"
	TokenLookup string

	//? TokenHeadName is a string in the header. Default value is "Bearer"
	TokenHeadName string

	//? TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
	TimeFunc func() time.Time

	//? HTTP Status messages for when something in the JWT middleware fails.
	//? Check error (e) to determine the appropriate error message.
	HTTPStatusMessageFunc func(e error, c *gin.Context) string

	//? Optionally return the token as a cookie
	SendCookie bool

	//? Allow insecure cookies for development over http
	SecureCookie bool

	//? Allow cookies to be accessed client side for development
	CookieHTTPOnly bool

	//? Allow cookie domain change for development
	CookieDomain string

	//? SendAuthorization allow return authorization header for every request
	SendAuthorization bool

	//? Disable abort() of context.
	DisabledAbort bool

	//? CookieName allow cookie name change for development
	CookieName string
}

//? New for check error with GinJWTMiddleware
func New(m *GinJWTMiddleware) (*GinJWTMiddleware, error) {
	if err := m.MiddlewareInit(); err != nil {
		return nil, err
	}

	return m, nil
}

func (mw *GinJWTMiddleware) MiddlewareInit() error {

	if mw.TokenLookup == "" {
		mw.TokenLookup = "header:Authorization"
	}

	if mw.SigningAlgorithm == "" {
		mw.SigningAlgorithm = "HS256"
	}

	if mw.TimeFunc == nil {
		mw.TimeFunc = time.Now
	}

	mw.TokenHeadName = strings.TrimSpace(mw.TokenHeadName)
	if len(mw.TokenHeadName) == 0 {
		mw.TokenHeadName = "Bearer"
	}

	if mw.Authorizator == nil {
		mw.Authorizator = func(data interface{}, c *gin.Context) bool {
			return true
		}
	}

	if mw.Unauthorized == nil {
		mw.Unauthorized = func(c *gin.Context, code int, message string) {
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": message,
			})
		}
	}

	if mw.LoginResponse == nil {
		mw.LoginResponse = func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, gin.H{
				"code":   http.StatusOK,
				"token":  token,
				"expire": expire.Format(time.RFC3339),
			})
		}
	}

	if mw.RefreshResponse == nil {
		mw.RefreshResponse = func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, gin.H{
				"code":   http.StatusOK,
				"token":  token,
				"expire": expire.Format(time.RFC3339),
			})
		}
	}

	if mw.IdentityKey == "" {
		mw.IdentityKey = "identity"
	}

	if mw.IdentityHandler == nil {
		mw.IdentityHandler = func(c *gin.Context) interface{} {
			claims := ExtractClaims(c)
			return claims
		}
	}

	if mw.HTTPStatusMessageFunc == nil {
		mw.HTTPStatusMessageFunc = func(e error, c *gin.Context) string {
			return e.Error()
		}
	}

	if mw.Realm == "" {
		mw.Realm = "gin jwt"
	}

	if mw.CookieName == "" {
		mw.CookieName = "jwt"
	}

	return nil
}

/// 																		//

//? ExtractClaims help to extract the JWT claims
func ExtractClaims(c *gin.Context) MapClaims {
	claims, exists := c.Get(JwtPayloadKey)
	if !exists {
		return make(MapClaims)
	}

	return claims.(MapClaims)
}

//? ExtractClaimsFromToken help to extract the JWT claims from token
func ExtractClaimsFromToken(token *jwt.Token) MapClaims {
	if token == nil {
		return make(MapClaims)
	}

	claims := MapClaims{}
	for key, value := range token.Claims.(jwt.MapClaims) {
		claims[key] = value
	}

	return claims
}

//? GetToken help to get the JWT token string
func GetToken(c *gin.Context) string {
	token, exists := c.Get("JWT_TOKEN")
	if !exists {
		return ""
	}

	return token.(string)
}
