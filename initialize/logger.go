package initialize

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/xsolare/re-chinese-go-backend/global"
)

func Logger() {
	global.GV_LOG = logrus.New()
	global.GV_LOG.SetOutput(os.Stdout)
	// global.GV_LOG.SetFormatter(&logrus.JSONFormatter{})
}
