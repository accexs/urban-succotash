package main

import (
	"log"
	"me-wallet/src/utils"
)

func main() {
	err := utils.CreateServer().Run()
	if err != nil {
		log.Panicln(err)
	}
}
