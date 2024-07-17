package examples

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func ShowGracefulRestartOrStop() error {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome to Gin Server")
	})

	srv := &http.Server{
		Addr:    ":8094",
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err.Error())
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with a timeout of 5 seconds
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't ba catch, so don't need to add it.
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(), timeout of 5 seconds
	select {
	// will block until the ctx.Done()
	case <-ctx.Done():
		log.Println("timeout of 5 seconds")
	}
	log.Println("Server existing")

	return nil
}
