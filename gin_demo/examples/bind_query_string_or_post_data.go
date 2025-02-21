package examples

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func ShowBindQueryStringOrPostData() error {
	r := gin.Default()

	r.GET("/testing", startPage)
	r.POST("/testing", startPage)

	return r.Run(":8084")
}

func startPage(ctx *gin.Context) {
	var person Person

	// If `GET`, only `Form` binding engine (`query`) used.
	// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
	// See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
	// curl -X GET "localhost:8084/testing?name=appleboy&address=xyz&birthday=1992-03-15"
	// response: {"address":"xyz","birthday":"1992-03-15T00:00:00Z","name":"appleboy"}
	// curl --location "localhost:8084/testing" --header 'Content-Type: application/json' --data '{"name":"appleboy","address":"xyz","birthday":"1992-03-15T00:00:00Z"}'
	// response: {"address":"xyz","birthday":"1992-03-15T00:00:00Z","name":"appleboy"}
	if ctx.ShouldBind(&person) == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"name":     person.Name,
			"address":  person.Address,
			"birthday": person.Birthday,
		})
	} else {
		ctx.String(http.StatusOK, "Success")
	}
}
