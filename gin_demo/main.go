package main

import (
	"github.com/denglitong/golang/gin_demo/examples"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"log"
)

var (
	g errgroup.Group
)

func main() {
	// Using sync/errgroup + router to listen on multiple ports for multiple routes
	g.Go(showQuickStart)
	g.Go(examples.ShowAsciiJSON)
	g.Go(examples.ShowBindDataRequestWithCustomStruct)
	g.Go(examples.ShowBindHtmlCheckBoxes)
	g.Go(examples.ShowBindQueryStringOrPostData)
	g.Go(examples.ShowBindUri)
	g.Go(examples.ShowBuildSingleBinaryWithAssertTemplate)
	g.Go(examples.ShowControlLogOutputColor)
	g.Go(examples.ShowCustomHttpConfiguration)
	g.Go(examples.ShowCustomLogFile)
	g.Go(examples.ShowCustomMiddleware)
	g.Go(examples.ShowCustomValidators)
	g.Go(examples.ShowDefineRoutesLogFormat)
	g.Go(examples.ShowGoroutinesInsideAMiddleware)
	// g.Go(examples.ShowGracefulRestartOrStop)
	g.Go(examples.ShowGroupingRoutes)
	g.Go(examples.ShowHowToWriteLogFile)
	g.Go(examples.ShowHtmlRendering)
	g.Go(examples.ShowJsonPadding)
	g.Go(examples.ShowMapAsQueryStringOrPostFormParams)

	// TODO
	// Model biding and validation
	// Multipart/Urlencoded binding
	// Multipart/Urlencoded form
	// Multi template
	// Only bind query string
	// Parameters in path
	// PureJSON
	// Query and post form
	// Query string params
	// Redirects
	// Run multiple service
	// SecureJSON
	// Serving data from reader
	// Serving static files
	g.Go(examples.ShowServingStaticFiles)

	// Set and get a cookie
	// Support Let's Encrypt
	// Try to bind body into different structs
	// Upload files
	g.Go(examples.ShowUploadFile)

	// Using BasicAuth middleware
	// Using HTTP method
	// Using middleware
	// Without middleware by default
	// XML/JSON/YAML/ProtoBuf rendering

	if err := g.Wait(); err != nil {
		log.Fatalln(err)
	}
}

func showQuickStart() error {
	router := gin.Default()
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// listen and serve on 0.0.0.0:8080
	return router.Run(":8080")
}
