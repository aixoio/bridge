package main

import (
	"log"

	"github.com/aixoio/bridge/server/db"
	"github.com/aixoio/bridge/server/env"
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

}
