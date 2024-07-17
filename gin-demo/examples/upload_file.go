package examples

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var uploadDir = flag.String("upload_dir", "./upload", "dir to upload files")

// go run . --upload_dir=/Users/litong.deng/Downloads
func ShowUploadFile() error {
	r := gin.New()

	// Set a lower memory limit for multipart forms (default is 32MB)
	r.MaxMultipartMemory = 8 << 20 // 8MB

	// curl -X POST http://localhost:8101/upload \
	//  -F "upload[]=@/Users/litong.deng/Downloads/aws-lambda-functions.png" \
	//  -F "upload[]=@/Users/litong.deng/Downloads/aws-textract-project.png"
	r.POST("/upload", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		files := form.File["upload[]"]
		for _, file := range files {
			log.Println(file.Filename)
			dst := fmt.Sprintf("%s/%s", *uploadDir, file.Filename)
			if err := c.SaveUploadedFile(file, dst); err != nil {
				c.String(http.StatusInternalServerError, err.Error())
			}
		}
	})

	return r.Run(":8101")
}
