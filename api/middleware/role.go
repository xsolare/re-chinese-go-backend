package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/xsolare/re-chinese-go-backend/api/models"
	"github.com/xsolare/re-chinese-go-backend/api/service"
	"github.com/xsolare/re-chinese-go-backend/utils"
	"github.com/xsolare/re-chinese-go-backend/utils/jwtauth/jwtuser"
	"github.com/xsolare/re-chinese-go-backend/utils/response"
)

var userService = service.UserService{}

func Role(allowedRoles []int) gin.HandlerFunc {
	return func(c *gin.Context) {

		err, roles := userService.RoleById(jwtuser.GetUserId(c))

		if err != nil {
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}

		isRoleExist := false
		roleIds := utils.Map(roles, func(x models.Role) int { return x.Id })

		for _, role := range allowedRoles {
			if utils.Сontains(roleIds, role) {
				isRoleExist = true
				break
			}
		}

		if !isRoleExist {
			response.FailWithDetailed(gin.H{"reload": true}, "You dont have aсcess", c)
			c.Abort()
			return
		}

		c.Set("roles", roles)
		c.Next()
	}
}
