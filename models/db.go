package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq" //PostgreSQL Driver

	env "github.com/xsolare/re-chinese-go-backend/config"
)

var ormObject orm.Ormer

// ConnectToDb - Initializes the ORM and Connection to the postgres DB
func ConnectToDb() {
    orm.RegisterDriver("postgres", orm.DRPostgres)
    orm.RegisterDataBase("default", "postgres", fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", env.POSTGRES_USER,env.POSTGRES_PASSWORD,env.POSTGRES_DB,env.POSTGRES_HOST))
    orm.RegisterModel(new(Users))
    ormObject = orm.NewOrm()
}

// GetOrmObject - Getter function for the ORM object with which we can query the database
func GetOrmObject() orm.Ormer {
    return ormObject
}