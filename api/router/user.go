package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xsolare/re-chinese-go-backend/api/controllers"
	"github.com/xsolare/re-chinese-go-backend/api/middleware"
)

type UserRouter struct{}

/// 																	///

func (r *UserRouter) InitUserRouter(rg *gin.RouterGroup) {
	userRouter := rg.Group("/user").Use(middleware.OperationRecord())
	api := new(controllers.UserController)

	// .Use(middleware.Role([]int{models.ADMIN, models.PREMIUM}))

	userRouter.GET("/", api.GetAll)
	userRouter.POST("/sign-in", api.SignIn)
	userRouter.POST("/sign-up", api.SignUp)
	userRouter.Use(middleware.Auth()).POST("/auth", api.Auth)
}
