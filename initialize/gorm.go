package initialize

import (
	"os"

	"github.com/xsolare/re-chinese-go-backend/global"
	"github.com/xsolare/re-chinese-go-backend/api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	p := global.GV_CONFIG.Pgsql
	if p.Dbname == "" {
		return nil
	}
	pgsqlConfig := postgres.Config{
		DSN:                  p.Dsn(),
		PreferSimpleProtocol: false,
	}
	if db, err := gorm.Open(postgres.New(pgsqlConfig)); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(p.MaxIdleConns)
		sqlDB.SetMaxOpenConns(p.MaxOpenConns)
		return db
	}
}

func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		models.User{},
		models.Role{},
		models.Initial{},
		models.Final{},
		models.FinalTone{},
		models.Pinyin{},
		models.Language{},
		models.PartOfSpeech{},
		models.Hieroglyph{},
	)
	if err != nil {
		global.GV_LOG.Error("register table failed")
		os.Exit(0)
	}
	global.GV_LOG.Info("register table success")
}

func RunMigration(db *gorm.DB) {
	models.InitDb(global.GV_DB)
}
