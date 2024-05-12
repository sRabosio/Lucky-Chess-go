package main

import (
	"encoding/json"
	"fmt"
	"log"
	"luckyChess/controllers/board"
	"luckyChess/controllers/game"
	"luckyChess/controllers/index"
	"luckyChess/entities"
	GameStoreService "luckyChess/services/store"
	"os"

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
	startingSet := getStartingSet()

	router := gin.Default()

	store := cookie.NewStore([]byte("gameState"))

	router.Use(sessions.Sessions("gameState", store))

	gameStoreService.NewGame(*startingSet)

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
	board.Register(router, gameStoreService)
	//end register routes
}

func getStartingSet() *entities.BoardTemplate {
	bytes, err := os.ReadFile("gameTemplates/default.json")

	defer func() {
		if err == nil {
			return
		}
		panic(err)
	}()

	jsonOut := entities.BoardTemplate{}

	json.Unmarshal(bytes, &jsonOut.Template)

	return &jsonOut
}
