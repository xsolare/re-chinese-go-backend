package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xsolare/re-chinese-go-backend/api/controllers"
)

type UserRouter struct{}

/// 																	///

func (r *UserRouter) InitUserRouter(rg *gin.RouterGroup) {
	userRouter := rg.Group("/user")
	api := new(controllers.UserController)

	userRouter.GET("/", api.GetAll)
}
