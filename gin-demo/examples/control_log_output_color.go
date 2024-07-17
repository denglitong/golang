package examples

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ShowControlLogOutputColor By default, logs output on console should be colorized depending on
// the detected TTY.
func ShowControlLogOutputColor() error {
	// Never colorize logs
	gin.DisableConsoleColor()
	// Always colorize logs
	// gin.ForceConsoleColor()

	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	return router.Run(":8087")
}
