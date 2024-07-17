package examples

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// ShowCustomHttpConfiguration Use http.ListenAndServe() directly.
func ShowCustomHttpConfiguration() error {
	// router := gin.Default()
	// http.ListenAndServe(":8088", router)

	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	server := &http.Server{
		Addr:           ":8088",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return server.ListenAndServe()
}
