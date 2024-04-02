package game

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	r := router.Group("/game")

	r.GET("", getGame)
}

func getGame(context *gin.Context) {
	templ := template.Must(template.ParseFiles("./pages/game/game.html"))
	templ.Execute(context.Writer, nil)
}
