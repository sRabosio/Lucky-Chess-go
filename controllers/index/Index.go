package index

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const folderPath = "./templates/pages/"

func Register(router *gin.Engine) {
	r := router.Group("/")

	r.GET("", getIndex)
}

func getIndex(context *gin.Context) {

	context.HTML(
		http.StatusOK,
		"master.html",
		nil,
	)
}
