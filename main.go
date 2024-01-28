package main

import (
	"log"
	"sso/utils"
)

func init() {
	utils.ReadEnv()
}

func main() {
	error := RunServer()
	if error != nil {
		log.Fatal("Error occured while starting server.")
	}
}
