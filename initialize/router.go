package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/xsolare/re-chinese-go-backend/api/router"
)

type Routes struct {
	router *gin.Engine
}

//* Create groups route
func NewRoutes() Routes {
	r := Routes{
		router: gin.Default(),
	}

	api := r.router.Group("/api")

	routes := new(router.RouterGroup)

	routes.FinalRouter.InitFinalRouter(api)
	routes.UserRouter.InitUserRouter(api)

	return r
}

// //* Combine handlers
// func (r Routes) AddUsers(rg *gin.RouterGroup) {
// 	users := rg.Group("/users")

// 	users.GET("/", controllers.GetUsers)
// }

//* Startup
func (r Routes) Run(addr ...string) error {
	return r.router.Run()
}
