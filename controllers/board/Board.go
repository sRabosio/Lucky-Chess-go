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
	r.GET("/movePiece", movePiece)

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
			_gameTemplateService.GetTemplate("default"),
		)
	} else {
		game = _gameStoreService.GetGame(val.(string))
	}

	context.HTML(
		http.StatusOK,
		"board.html",
		entities.BoardViewState{
			Game: game,
		},
	)
}

func movePiece(context *gin.Context) {
	var err error
	game := _gameStoreService.GetGame("1")
	status := http.StatusOK

	defer func() {
		if err != nil {
			println("MovePiece -> " + err.Error())
			context.Error(err)
		}

		context.HTML(
			status,
			"board.html",
			entities.BoardViewState{
				Game: game,
			},
		)
	}()

	pieceX, err := strconv.Atoi(context.Query("pieceX"))

	if err != nil {
		status = http.StatusBadRequest
		return
	}

	pieceY, err := strconv.Atoi(context.Query("pieceY"))

	if err != nil {
		status = http.StatusBadRequest
		return
	}

	targetX, err := strconv.Atoi(context.Query("targetX"))

	if err != nil {
		status = http.StatusBadRequest
		return
	}

	targetY, err := strconv.Atoi(context.Query("targetY"))

	if err != nil {
		status = http.StatusBadRequest
		return
	}

	err = _gameStateService.MovePiece(
		&game,
		"1",
		entities.TileCoords{Row: pieceY, Tile: pieceX},
		entities.TileCoords{Row: targetY, Tile: targetX},
	)

	if err != nil {
		status = http.StatusBadRequest
		return
	}

	err = _gameStoreService.ApplyChanges("1", game)

	if err != nil {
		status = http.StatusBadRequest
		return
	}
}

func getMoves(context *gin.Context) {

	var err error
	var x, y int
	game := _gameStoreService.GetGame("1")
	status := http.StatusOK

	defer func() {
		if err != nil {
			println("GetMoves -> " + err.Error())
			context.Error(err)
			x = -1
			y = -1
		}

		context.HTML(
			status,
			"board.html",
			entities.BoardViewState{
				Game:      game,
				SelectedX: x,
				SelectedY: y,
			},
		)
	}()

	x, err = strconv.Atoi(context.Query("x"))

	if err != nil {
		status = http.StatusBadRequest
		return
	}

	y, err = strconv.Atoi(context.Query("y"))

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
