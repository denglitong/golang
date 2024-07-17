package examples

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PersonUri struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func ShowBindUri() error {
	r := gin.Default()

	r.GET("/:name/:id", bindUri)

	return r.Run(":8085")
}

// curl -v localhost:8085/thinkerou/987fbc97-4bed-5078-9f07-9141ba07c9f3
// 	200 {"name":"thinkerou","uuid":"987fbc97-4bed-5078-9f07-9141ba07c9f3"}
// curl -v localhost:8085/thinkerou/not-uuid
//  400 {"msg":"Key: 'PersonUri.ID' Error:Field validation for 'ID' failed on the 'uuid' tag"}
//  {"msg":[{}]}
func bindUri(c *gin.Context) {
	var person PersonUri
	if err := c.ShouldBindUri(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"name": person.Name, "uuid": person.ID})
}
