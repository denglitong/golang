package examples

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowGroupingRoutes() error {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/login", loginHandlerV1)
		v1.GET("/submit", submitHandlerV1)
		v1.GET("/read", readHandlerV1)
	}

	v2 := router.Group("/v2")
	{
		v2.GET("/login", loginHandlerV2)
		v2.GET("/submit", submitHandlerV2)
		v2.GET("/read", readHandlerV2)
	}

	return router.Run(":8095")
}

func loginHandlerV1(c *gin.Context) {
	c.String(http.StatusOK, "login.v1")
}

func submitHandlerV1(c *gin.Context) {
	c.String(http.StatusOK, "submit.v1")
}

func readHandlerV1(c *gin.Context) {
	c.String(http.StatusOK, "read.v1")
}

func loginHandlerV2(c *gin.Context) {
	c.String(http.StatusOK, "login.v2")
}

func submitHandlerV2(c *gin.Context) {
	c.String(http.StatusOK, "submit.v2")
}

func readHandlerV2(c *gin.Context) {
	c.String(http.StatusOK, "read.v2")
}
