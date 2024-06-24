package main

import (
	"fmt"

	"github.com/aixoio/bridge/server/env"
)

func main() {
	fmt.Println(env.LoadENV())
}
