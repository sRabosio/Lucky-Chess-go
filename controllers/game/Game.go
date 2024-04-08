package game

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

const folderPath = "./templates/pages/game/"

func Register(router *gin.Engine) {
	r := router.Group("/game")

	r.GET("", getGame)
}

func getGame(context *gin.Context) {
	templ := template.Must(template.ParseFiles(folderPath + "game.html"))
	templ.Execute(context.Writer, nil)
}
