package examples

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func ShowCustomMiddleware() error {
	r := gin.New()
	r.Use(Logger())

	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)
		log.Println(fmt.Sprintf("example: %v", example))
		c.JSON(http.StatusOK, gin.H{
			"example": example,
		})
	})

	return r.Run(":8090")
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		// set example variable
		c.Set("example", "12345")

		// before request
		c.Next()
		// after request
		latency := time.Since(t)
		log.Print(fmt.Sprintf("latency: %v", latency))
		// access the status we are sending
		status := c.Writer.Status()
		log.Println(fmt.Sprintf("status: %v", status))
	}
}
