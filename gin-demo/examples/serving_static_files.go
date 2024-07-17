package examples

import (
	"flag"
	"github.com/gin-gonic/gin"
	"net/http"
)

var staticDir = flag.String("static_dir", "./", "dir to serve static files")

// go run . --static_dir=/Users/litong.deng/Downloads
func ShowServingStaticFiles() error {
	flag.Parse()
	r := gin.New()
	// can reflect the newest files under the dir
	r.StaticFS("/", http.Dir(*staticDir))

	return r.Run(":8100")
}
