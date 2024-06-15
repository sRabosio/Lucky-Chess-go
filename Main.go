package main

import (
	"fmt"
	"log"
	"luckyChess/controllers/board"
	"luckyChess/controllers/game"
	"luckyChess/controllers/index"
	"luckyChess/services/gameStateService"
	"luckyChess/services/gameTemplateService"
	GameStoreService "luckyChess/services/store"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const prodPort = ""

// singleton services
var gameStoreService = GameStoreService.New()

func main() {
	fmt.Println("starting server")

	//todo: replace w/ port from args
	println("strating set")

	router := gin.Default()

	store := cookie.NewStore([]byte("gameCode"))
	router.Use(sessions.Sessions("gameCode", store))

	//register assets
	router.Static("static", "./assets")

	//load templates
	router.LoadHTMLGlob("templates/**/*.html")
	router.LoadHTMLGlob("templates/**/**/*.html")

	initRoutes(router)

	//start server
	log.Fatal(router.Run())
}

func initRoutes(router *gin.Engine) {

	//register routes
	index.Register(router)
	game.Register(router, gameStoreService)
	board.Register(router, gameStoreService,
		gameTemplateService.New(),
		gameStateService.New())
	//end register routes
}
