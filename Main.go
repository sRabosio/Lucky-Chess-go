package main

import (
	"fmt"
	"log"
	"luckyChess/controllers/board"
	"luckyChess/controllers/game"
	"luckyChess/controllers/index"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const prodPort = ""

func main() {
	fmt.Println("starting server")

	//todo: replace w/ port from args

	router := gin.Default()

	store := cookie.NewStore([]byte("gameState"))

	router.Use(sessions.Sessions("gameState", store))

	//register assets
	router.Static("static", "./assets")

	router.LoadHTMLGlob("templates/**/*.html")
	router.LoadHTMLGlob("templates/**/**/*.html")

	//register routes
	index.Register(router)
	game.Register(router)
	board.Register(router)
	//end register routes

	//start server
	log.Fatal(router.Run())
}
