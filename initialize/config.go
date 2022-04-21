package initialize

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	g "github.com/xsolare/re-chinese-go-backend/global"
)

func Config() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	g.GV_CONFIG.Pgsql.Host = os.Getenv("POSTGRES_HOST")
	g.GV_CONFIG.Pgsql.Port = os.Getenv("POSTGRES_PORT")
	g.GV_CONFIG.Pgsql.Username = os.Getenv("POSTGRES_USER")
	g.GV_CONFIG.Pgsql.Password = os.Getenv("POSTGRES_PASSWORD")
	g.GV_CONFIG.Pgsql.Dbname = os.Getenv("POSTGRES_DB")
	g.GV_CONFIG.Pgsql.MaxIdleConns = 8
	g.GV_CONFIG.Pgsql.MaxOpenConns = 8

	g.GV_CONFIG.JWT.SecretKey = os.Getenv("JWT_SECRET")
	g.GV_CONFIG.JWT.ExpiresTime = 10 * 24 * 60 * 60
	g.GV_CONFIG.JWT.BufferTime = 10 * 24 * 60 * 60
}
