package main

import (
	"log"
	"pos_api/config"
	"pos_api/jwt"
	"pos_api/router"
)

func main() {
	// setup config
	config.Setup()

	// setup jwt
	jwt.Setup()

	// init db
	config.InitDb()

	// init router
	r := router.InitRouter()

	// start server
	log.Fatal(r.Run())
}
