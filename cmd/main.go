package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/xsolare/re-chinese-go-backend/app"
	config "github.com/xsolare/re-chinese-go-backend/configs"
	"github.com/xsolare/re-chinese-go-backend/internal/models"

	"github.com/astaxie/beego/orm"
)

var ORM orm.Ormer

func init() {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&logrus.JSONFormatter{})

	config.GetEnvVariables()
	models.ConnectToDb()
	ORM = models.GetOrmObject()

}

func main() {
    r := app.NewRoutes()

	r.Run(":8080")
}
