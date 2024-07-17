package examples

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// ShowDefineRoutesLogFormat
//  The default log of routes is:
// 	[GIN-debug] POST /foo	--> main.main.func1 (3 handlers)
// Use gin.DebugPrintRouteFunc to custom this information in given format.
func ShowDefineRoutesLogFormat() error {
	r := gin.Default()

	// Global config, effect all gi.Default() routes even serve in another ports
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	r.POST("/foo", func(c *gin.Context) {
		c.String(http.StatusOK, "foo")
	})
	r.GET("/bar", func(c *gin.Context) {
		c.String(http.StatusOK, "BAR")
	})
	r.GET("/status", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	return r.Run(":8092")
}
