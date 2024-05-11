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

func main() {
	fmt.Println("starting server")

	//todo: replace w/ port from args
	println("strating set")
	getStartingSet()

	router := gin.Default()

	store := cookie.NewStore([]byte("gameState"))

	gameStoreService := GameStoreService.New()

	router.Use(sessions.Sessions("gameState", store))

	//register assets
	router.Static("static", "./assets")

	router.LoadHTMLGlob("templates/**/*.html")
	router.LoadHTMLGlob("templates/**/**/*.html")

	//register routes
	index.Register(router)
	game.Register(router, gameStoreService)
	board.Register(router)
	//end register routes

	//start server
	log.Fatal(router.Run())
}

func getStartingSet() *entities.Board {
	bytes, err := os.ReadFile("gameTemplates/default.json")

	defer func() {
		if err == nil {
			return
		}
		panic(err)
	}()

	var res *entities.Board

	json.Unmarshal(bytes, res)

	return res
}
