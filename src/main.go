package main

import (
	"github.com/shwatanap/workout-wizard-api/src/config"
	"github.com/shwatanap/workout-wizard-api/src/infra/db/driver"
	"github.com/shwatanap/workout-wizard-api/src/presen/router"
	"github.com/shwatanap/workout-wizard-api/src/wire"
)

func main() {
	config.LoadEnv()
	db := driver.NewDriver()
	menuHandler := wire.InitMenuHandler(db)

	router := router.InitRouting(menuHandler)
	router.Run()
}
