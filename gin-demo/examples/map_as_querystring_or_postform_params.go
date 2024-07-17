package examples

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowMapAsQueryStringOrPostFormParams() error {
	r := gin.New()

	// POST /post?ids[a]=1234&ids[b]=hello HTTP/1.1
	// Content-Type: application/x-www-form-urlencoded
	//
	// names[first]=thinkerou&names[second]=tianou
	// curl --location 'http://localhost:8098/post?ids[a]=1234&ids[b]=hello' \
	// --header 'Content-Type: application/x-www-form-urlencoded' \
	// --data-urlencode 'names%5Bfirst%5D=thinkerou' \
	// --data-urlencode 'names%5Bsecon%5D=tianou'
	r.POST("/post", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		c.JSONP(http.StatusOK, gin.H{
			"ids":   ids,
			"names": names,
		})
	})

	return r.Run(":8099")
}
