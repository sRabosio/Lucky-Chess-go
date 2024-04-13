package board

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ChessBoardData struct {
	Rows [8]ChessBoardRow
}

type ChessBoardRow struct {
	Tiles [8]ChessBoardTile
}

type ChessBoardTile struct {
	Content string
}

func Register(router *gin.Engine) {

	r := router.Group("/board")
	r.GET("", getBoard)
}

func getBoard(context *gin.Context) {

	data := ChessBoardData{}
	for i := 0; i < 8; i++ {
		row := ChessBoardRow{}
		for j := 0; j < 8; j++ {
			tile := ChessBoardTile{
				Content: strconv.Itoa(j),
			}
			row.Tiles[j] = tile
		}
		data.Rows[i] = row
	}

	context.HTML(
		http.StatusOK,
		"board.html",
		data,
	)
}
