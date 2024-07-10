package index

import (
	eSessionStatus "luckyChess/entities/ESessionStatus"
	"luckyChess/services/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

const folderPath = "./templates/pages/"

var _gameStoreService interfaces.IGameStoreService
var _userService interfaces.IUserService

func Register(
	router *gin.Engine,
	userService interfaces.IUserService,
	gameStoreService interfaces.IGameStoreService,
) {

	_userService = userService
	_gameStoreService = gameStoreService

	r := router.Group("/")

	r.GET("", getIndex)
}

func getSessionStatus(req *http.Request) (eSessionStatus.ESessionStatus, string, string) {

	userCodeCookie, err := req.Cookie("usercode")
	if err != nil {
		return eSessionStatus.NO_USER, "", ""
	}

	hasUser, err := _userService.HasUser(userCodeCookie.Value)
	if err != nil || !hasUser {
		return eSessionStatus.NO_USER, "", ""
	}

	gameCodeCookie, err := req.Cookie("gamecode")
	if err != nil {
		return eSessionStatus.NO_GAME, userCodeCookie.Value, ""
	}
	_, err = _gameStoreService.GetGame(gameCodeCookie.Value)
	if err != nil {
		return eSessionStatus.NO_GAME, userCodeCookie.Value, ""
	}

	return eSessionStatus.IN_GAME, userCodeCookie.Value, gameCodeCookie.Value
}

func getIndex(context *gin.Context) {

	sessionStatus, _, _ := getSessionStatus(context.Request)

	context.HTML(
		http.StatusOK,
		"master.html",
		gin.H{
			"sessionStatus": sessionStatus.String(),
		},
	)
}
