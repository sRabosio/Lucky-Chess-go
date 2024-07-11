package userController

import (
	"luckyChess/services/interfaces"
	"net/http"
	"strings"

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

	r := router.Group("/user")

	r.POST("registerNickname", registerNickname)
}

func registerNickname(context *gin.Context) {
	nickname := context.PostForm("nickname")

	if strings.Trim(nickname, " ") == "" {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user, err := _userService.GenerateUser(nickname)

	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	context.SetCookie("usercode", user.Code, 500000, "/", "localhost", true, true)

	context.Status(http.StatusCreated)
}
