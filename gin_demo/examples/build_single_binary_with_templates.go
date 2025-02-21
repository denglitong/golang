package examples

import (
	"github.com/denglitong/golang/gin_demo/templates"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Prepare Packages
// 	go get -u github.com/jessevdk/go-assets-builder
// 	go install github.com/jessevdk/go-assets-builder
// Generate asserts.go:
// 	cd examples && go-assets-builder -p examples ./templates -o ./asserts.go
// Build the single file server
// 	go build -o assets-in-binary
// Run the single file server
// 	./assets-in-binary

func ShowBuildSingleBinaryWithAssertTemplate() error {
	r := gin.New()
	t, err := templates.LoadTemplates()
	if err != nil {
		panic(err)
	}
	r.SetHTMLTemplate(t)
	r.GET("/", serveIndexPage)
	r.GET("/bar", serveBarPage)

	return r.Run(":8086")
}

func serveIndexPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "/index.htm", gin.H{
		"Foo": "World",
	})
}

func serveBarPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "/bar.htm", gin.H{
		"Bar": "World",
	})
}
