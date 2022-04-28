package initialize

import (
	"fmt"

	"github.com/xsolare/re-chinese-go-backend/global"
	"github.com/xsolare/re-chinese-go-backend/storage/redis"
	"github.com/xsolare/re-chinese-go-backend/utils"
)

func Redis() error {
	var err error = nil
	global.GV_REDIS, err = redis.NewRedis(global.GV_CONFIG.Redis)

	global.GV_REDIS.Process("FLUSHALL")

	fmt.Println(utils.Cyan("- Redis -"))

	return err
}
