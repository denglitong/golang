package examples

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// When starting new goroutines inside a middleware or handler, you SHOULD NOT use the
// original context inside it, you have to use a read-only copy.

func ShowGoroutinesInsideAMiddleware() error {
	r := gin.Default()

	r.GET("/long_async", func(c *gin.Context) {
		// create a copy to be used inside the goroutine
		cCp := c.Copy()
		go func() {
			time.Sleep(5 * time.Second)
			// note that you are using the copied context "cCp", IMPORTANT
			log.Println("Done! in path " + cCp.Request.URL.Path)
		}()
		c.String(http.StatusOK, "ok")
	})

	r.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		// since we are NOT using a goroutine, we do not have to copy the context
		log.Println("Done! in path " + c.Request.URL.Path)
		c.String(http.StatusOK, "ok")
	})

	return r.Run(":8093")
}
