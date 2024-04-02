package index

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	r := router.Group("/")

	r.GET("", getIndex)
}

func getIndex(context *gin.Context) {

	templ := template.Must(template.ParseFiles("./pages/master.html"))
	templ.Execute(context.Writer, nil)
}
