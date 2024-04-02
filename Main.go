package main

import (
	"fmt"
	"log"
	"luckyChess/controllers/game"
	"luckyChess/controllers/index"

	"github.com/gin-gonic/gin"
)

const prodPort = ""

func main() {
	fmt.Println("starting server")

	//todo: replace w/ port from args

	router := gin.Default()

	//register assets
	router.Static("static", "./assets")

	//register routes
	index.Register(router)
	game.Register(router)
	//end register routes

	//start server
	log.Fatal(router.Run())
}
