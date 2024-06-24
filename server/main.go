package main

import (
	"log"

	"github.com/aixoio/bridge/server/db"
	"github.com/aixoio/bridge/server/env"
	"github.com/aixoio/bridge/server/router"
)

func main() {
	dat, err := env.LoadENV()
	if err != nil {
		log.Panicln("Cannot load env")
	}

	err = db.Connect(dat)
	if err != nil {
		log.Panicln("Cannot connect to PostgreSQL")
	}
	defer db.Disconnect()

	router.StartServer(dat)

}
