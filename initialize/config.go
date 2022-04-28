package initialize

import (
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"github.com/xsolare/re-chinese-go-backend/configs"
	g "github.com/xsolare/re-chinese-go-backend/global"
	"github.com/xsolare/re-chinese-go-backend/utils"
)

func Config() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	//? Pgsql
	g.GV_CONFIG.Pgsql.Host = os.Getenv("POSTGRES_HOST")
	g.GV_CONFIG.Pgsql.Port = os.Getenv("POSTGRES_PORT")
	g.GV_CONFIG.Pgsql.Username = os.Getenv("POSTGRES_USER")
	g.GV_CONFIG.Pgsql.Password = os.Getenv("POSTGRES_PASSWORD")
	g.GV_CONFIG.Pgsql.Dbname = os.Getenv("POSTGRES_DB")
	g.GV_CONFIG.Pgsql.MaxIdleConns = 8
	g.GV_CONFIG.Pgsql.MaxOpenConns = 8

	//? Jwt
	g.GV_CONFIG.JWT.SecretKey = os.Getenv("JWT_SECRET")
	g.GV_CONFIG.JWT.ExpiresTime = 10 * 24 * 60 * 60
	g.GV_CONFIG.JWT.BufferTime = 10 * 24 * 60 * 60

	//? Cors
	g.GV_CONFIG.Cors.Mode = "whitelist"
	g.GV_CONFIG.Cors.Whitelist = []configs.CORSWhitelist{{
		AllowOrigin:      "localhost",
		AllowMethods:     "Content-Type,AccessToken,X-CSRF-Token,Authorization,Token,X-Token,X-User-Id,X-Auth",
		AllowHeaders:     "POST,GET,OPTIONS,DELETE,PUT",
		ExposeHeaders:    "Content-Length,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Content-Type",
		AllowCredentials: true,
	}}

	//? Redis
	g.GV_CONFIG.Redis = &redis.Options{}
	g.GV_CONFIG.Redis.Addr = os.Getenv("REDIS_ADDR")
	g.GV_CONFIG.Redis.Password = os.Getenv("REDIS_PASSWORD")
	g.GV_CONFIG.Redis.DB, _ = utils.StringToInt(os.Getenv("REDIS_DB"))

	fmt.Printf(utils.Cyan("- Env -"))

}
