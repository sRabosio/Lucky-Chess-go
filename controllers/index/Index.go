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
	r.GET("getSessionStatus", func(ctx *gin.Context) {
		sessionStatus, _, _ := getSessionStatus(ctx.Request)

		ctx.JSON(http.StatusOK, gin.H{
			"sessionStatus": sessionStatus,
		})
	})
	r.GET("getStatusDialog", getStatusDialog)
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

func getStatusDialog(context *gin.Context) {
	status, _, _ := getSessionStatus(context.Request)

	if status == eSessionStatus.IN_GAME {
		context.String(http.StatusOK, "")
		return
	}

	context.HTML(http.StatusOK, status.String()+".html", nil)
}
