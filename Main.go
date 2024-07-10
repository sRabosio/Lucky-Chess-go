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
	"luckyChess/services/userService"

	"github.com/gin-gonic/gin"
)

const prodPort = ""

// singleton services
var gameStoreService = GameStoreService.New()
var _userService = userService.New()

func main() {
	fmt.Println("starting server")

	//todo: replace w/ port from args

	router := gin.Default()

	//register assets
	router.Static("static", "./assets")

	//load templates
	router.LoadHTMLFiles(
		"templates/components/chessboard/board.html",
		"templates/pages/game/game.html",
		"templates/pages/master.html")
	initRoutes(router)

	//start server
	log.Fatal(router.Run())
}

func initRoutes(router *gin.Engine) {

	//register routes
	index.Register(router,
		_userService,
		gameStoreService)
	game.Register(router, gameStoreService)
	board.Register(router, gameStoreService,
		gameTemplateService.New(),
		gameStateService.New())
	//end register routes
}
