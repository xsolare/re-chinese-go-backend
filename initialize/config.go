package initialize

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/xsolare/re-chinese-go-backend/global"
)

func Config() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	global.GV_CONFIG.Pgsql.Host = os.Getenv("POSTGRES_HOST")
	global.GV_CONFIG.Pgsql.Port = os.Getenv("POSTGRES_PORT")
	global.GV_CONFIG.Pgsql.Username = os.Getenv("POSTGRES_USER")
	global.GV_CONFIG.Pgsql.Password = os.Getenv("POSTGRES_PASSWORD")
	global.GV_CONFIG.Pgsql.Dbname = os.Getenv("POSTGRES_DB")
	global.GV_CONFIG.Pgsql.MaxIdleConns = 8
	global.GV_CONFIG.Pgsql.MaxOpenConns = 8
}
