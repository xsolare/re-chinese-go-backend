package global

import (
	"github.com/sirupsen/logrus"
	"github.com/xsolare/re-chinese-go-backend/configs"

	"gorm.io/gorm"
)

var (
	GV_DB     *gorm.DB
	GV_CONFIG configs.Server
	GV_LOG    *logrus.Logger
)
