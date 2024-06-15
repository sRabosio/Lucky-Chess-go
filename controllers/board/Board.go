package board

import (
	"luckyChess/entities"
	"luckyChess/services/interfaces"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var _gameStoreService interfaces.IGameStoreService
var _gameTemplateService interfaces.IGameTemplatesService
var _gameStateService interfaces.IGameStateService

func Register(router *gin.Engine,
	gameStoreService interfaces.IGameStoreService,
	gameTemplateService interfaces.IGameTemplatesService,
	gameStateService interfaces.IGameStateService) {

	r := router.Group("/board")
	r.GET("", getBoard)
	r.GET("/getMoves", getMoves)

	_gameStoreService = gameStoreService
	_gameTemplateService = gameTemplateService
	_gameStateService = gameStateService
}

func getBoard(context *gin.Context) {

	store, err := cookie.Store.Get(cookie.NewStore(), context.Request, "gameCode")

	var game entities.Game

	if err != nil {
		context.AbortWithError(500, err)
		return
	}

	val := store.Values["code"]

	if val == nil {
		game = _gameStoreService.NewGame(
			_gameTemplateService.GetTemplate("debug_nopawns"),
		)
	} else {
		game = _gameStoreService.GetGame(val.(string))
	}

	context.HTML(
		http.StatusOK,
		"board.html",
		game,
	)
}

func getMoves(context *gin.Context) {

	var err error
	game := _gameStoreService.GetGame("1")
	status := http.StatusOK

	defer func() {
		if err != nil {
			println("GetMoves -> " + err.Error())
			context.Error(err)
		}

		context.HTML(
			status,
			"board.html",
			game,
		)
	}()

	x, err := strconv.Atoi(context.Query("x"))

	if err != nil {
		status = http.StatusBadRequest
		return
	}

	y, err := strconv.Atoi(context.Query("y"))

	if err != nil {
		status = http.StatusBadRequest
		return
	}

	moveSet, err := _gameStateService.GetMoveset(&game, "1", entities.TileCoords{Tile: x, Row: y})

	if err != nil {
		status = http.StatusOK
		return
	}

	selectedTile := &game.Board.Rows[y].Tiles[x]

	//selected
	if selectedTile.Piece > 0 {
		selectedTile.State = "selected"
	}

	for _, move := range moveSet {
		currentTile := &game.Board.Rows[move.Row].Tiles[move.Tile]

		if currentTile.Piece > 0 {
			currentTile.State = "eat"
		} else {
			currentTile.State = "highlighted"
		}
	}
}
