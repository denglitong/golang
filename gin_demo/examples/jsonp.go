package examples

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ShowJsonPadding Solve Ajax not allowed to fetch cross-domain data issue.
func ShowJsonPadding() error {
	r := gin.New()

	r.GET("/JSONP", func(c *gin.Context) {
		fetchedAnotherServerData := gin.H{
			"foo": "bar",
		}
		// curl -X GET localhost:8098/JSONP?callback=jsCallbackFunc
		// return: jsCallbackFunc({"foo":"bar"});
		//
		// usage: <script src="http://example.com/JSONP?callback=handleResponse"></script>
		// equivalent to:<script>jsCallbackFunc({"foo":"bar"});</script>
		c.JSONP(http.StatusOK, fetchedAnotherServerData)
	})

	return r.Run(":8098")
}
