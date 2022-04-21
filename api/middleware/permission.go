package middleware

import (
	"github.com/gin-gonic/gin"
)

func AuthCheckRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		// data, _ := c.Get(jwt.JwtPayloadKey)

		// fmt.Println("data")
		// fmt.Println(data)

		// v := data.(jwt.MapClaims)

		// fmt.Println("v")
		// fmt.Println(v)

		// if v["rolekey"] == "admin" {
		// 	c.Next()
		// 	return
		// } else {
		// 	c.JSON(http.StatusOK, gin.H{
		// 		"code": 403,
		// 		"msg":  "Nope",
		// 	})
		// 	c.Abort()
		// 	return
		// }
	}
}
