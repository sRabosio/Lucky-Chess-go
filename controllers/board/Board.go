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

func Register(router *gin.Engine, gameStoreService interfaces.IGameStoreService, gameTemplateService interfaces.IGameTemplatesService) {

	r := router.Group("/board")
	r.GET("", getBoard)
	r.GET("/getMoves", getMoves)

	_gameStoreService = gameStoreService
	_gameTemplateService = gameTemplateService
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
		game,
	)
}

func getMoves(context *gin.Context) {

	game := _gameStoreService.GetGame("1")
	status := http.StatusOK

	defer func() {
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

	//selected
	selectedTile := &game.Board.Rows[y].Tiles[x]

	if selectedTile.Piece == 0 {
		return
	}

	selectedTile.State = "selected"
	//highlighted
	game.Board.Rows[y+1].Tiles[x].State = "highlighted"

}
