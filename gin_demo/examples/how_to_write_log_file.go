package examples

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func ShowHowToWriteLogFile() error {
	gin.DisableConsoleColor()

	f, _ := os.Create("gin.log")
	// Write the logs to file and also the console at the same time.
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return router.Run(":8096")
}
