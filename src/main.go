package main

import (
	"log"
	"me-wallet/src/helpers"
)

func main() {
	err := helpers.CreateServer().Run()
	if err != nil {
		log.Panicln(err)
	}
}
