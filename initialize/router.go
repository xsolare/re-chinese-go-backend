package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
	"github.com/xsolare/re-chinese-go-backend/api/middleware"
	"github.com/xsolare/re-chinese-go-backend/api/router"
	"github.com/xsolare/re-chinese-go-backend/global"
)

type Routes struct {
	router *gin.Engine
}

//* Create groups route
func NewRoutes() Routes {
	r := Routes{
		router: gin.Default(),
	}

	//? Swagger
	// r.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//? Cors
	r.router.Use(middleware.Cors())
	// r.router.Use(middleware.CorsByRules())

	r.router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api := r.router.Group("/api/")

	routes := new(router.RouterGroup)

	routes.FinalRouter.InitFinalRouter(api)
	routes.UserRouter.InitUserRouter(api)
	routes.WordRouter.InitWordRouter(api)
	routes.PinyinRouter.InitPinyinRouter(api)
	routes.WordCollectionRouter.InitWordCollectionRouter(api)
	routes.HieroglyphCollectionRouter.InitHieroglyphCollectionRouter(api)
	routes.HieroglyphRouter.InitHieroglyphRouter(api)

	return r
}

//* Startup
func (r Routes) Run() error {
	return r.router.Run(global.GV_CONFIG.Gin.Port)
}

func (r Routes) RunTLS() error {
	r.router.Use(LoadTLS())
	return r.router.RunTLS(global.GV_CONFIG.Gin.Port, global.GV_CONFIG.Gin.Cert, global.GV_CONFIG.Gin.Key)
}

func LoadTLS() gin.HandlerFunc {
	return func(c *gin.Context) {
		middleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     global.GV_CONFIG.Gin.Host + global.GV_CONFIG.Gin.Port,
		})
		err := middleware.Process(c.Writer, c.Request)
		if err != nil {
			global.GV_LOG.Error(err)
			return
		}
		c.Next()
	}
}
