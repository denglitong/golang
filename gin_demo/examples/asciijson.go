package examples

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowAsciiJSON() error {
	r := gin.Default()

	r.GET("/someJSON", func(context *gin.Context) {
		data := map[string]interface{}{
			"lang": "Go语言",
			"tag":  "<br/>",
		}
		context.AsciiJSON(http.StatusOK, data)
	})

	return r.Run(":8081")
}
