package main

import (
	"github.com/xsolare/re-chinese-go-backend/initialize"
)

func main() {
	initialize.Config()

	//?		Init global variables			?//
	initialize.Logger()
	initialize.Gorm()
	initialize.Redis()

	//?		Init model / Run migration		?//
	// initialize.ResetSchema(global.GV_DB)
	// initialize.RegisterTables(global.GV_DB)
	// initialize.RunMigration(global.GV_DB)

	//?		Gin routes and run server 		?//
	r := initialize.NewRoutes()
	r.RunTLS()
}
