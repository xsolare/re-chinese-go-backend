package main

import (
	"github.com/xsolare/re-chinese-go-backend/global"
	"github.com/xsolare/re-chinese-go-backend/initialize"
)

func main() {
	initialize.Config()

	//?		Init global variables			?//
	global.GV_LOG = initialize.Logger()
	global.GV_DB = initialize.Gorm()

	//?		Init model / Run migration		?//
	// initialize.ResetSchema(global.GV_DB)
	// initialize.RegisterTables(global.GV_DB)
	// initialize.RunMigration(global.GV_DB

	//?		Gin routes and run server 		?//
	r := initialize.NewRoutes()
	r.Run(":8080")
}
