package main

import (
	"gin-api-template/internal/env"
	"gin-api-template/internal/server"
	"log"
)

func main() {
	env, err := env.NewValue()
	if err != nil {
		log.Fatal(err)
	}

	server.Run(env)
}
