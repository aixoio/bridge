package main

import (
	"fmt"
	"log"

	"github.com/aixoio/bridge/server/env"
)

func main() {
	dat, err := env.LoadENV()
	if err != nil {
		log.Fatalln("Cannot load env")
	}

	fmt.Println(dat)
}
