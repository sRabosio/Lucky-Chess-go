package board

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

type ChessBoardData struct {
	rows string
}

const boardFolder = "./templates/components/chessboard/"

func Register(router *gin.Engine) {
	r := router.Group("/board")

	r.GET("", getBoard)
}

func getBoard(context *gin.Context) {
	templ := template.Must(template.ParseFiles(boardFolder + "board.html"))

	data := ChessBoardData{
		rows: "",
	}

	templ.Execute(context.Writer, data)
}
