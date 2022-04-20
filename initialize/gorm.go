package initialize

import (
	"os"

	"github.com/xsolare/re-chinese-go-backend/api/models"
	"github.com/xsolare/re-chinese-go-backend/global"
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
		models.Role{},
		models.User{},
		models.UserStat{},
		models.Initial{},
		models.Final{},
		models.FinalTone{},
		models.Pinyin{},
		models.Language{},
		models.PartOfSpeech{},
		models.Categorie{},
		models.Hieroglyph{},
		models.HieroglyphCollection{},
		models.HieroglyphTranslate{},
		models.HieroglyphKey{},
		models.Word{},
		models.WordHieroglyph{},
		models.WordTranslate{},
		models.WordCollection{},
		models.OperationRecord{},
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

func ResetSchema(db *gorm.DB) {
	db.Raw("drop schema public cascade;")
	db.Raw("create schema public;")
}
