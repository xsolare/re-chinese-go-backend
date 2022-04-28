package global

import (
	"github.com/sirupsen/logrus"
	"github.com/xsolare/re-chinese-go-backend/configs"
	"github.com/xsolare/re-chinese-go-backend/storage/redis"

	"gorm.io/gorm"
)

var (
	GV_DB     *gorm.DB
	GV_REDIS  *redis.Redis
	GV_LOG    *logrus.Logger
	GV_CONFIG configs.Server
)
