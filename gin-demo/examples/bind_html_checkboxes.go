package examples

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type MyForm struct {
	Name   string   `form:"name"`
	Colors []string `form:"colors[]"`
}

func handleForm(ctx *gin.Context) {
	var form MyForm
	err := ctx.ShouldBind(&form)
	if err != nil {
		log.Fatalln(err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"name":   form.Name,
		"colors": form.Colors,
	})
}

func ShowBindHtmlCheckBoxes() error {
	r := gin.Default()

	// curl --location 'http://localhost:8083/form' \
	// --header 'Content-Type: application/json' \
	// --data '{"name":"go","colors":["red","green","blue"]}'
	// return: {"colors":["red","green","blue"],"name":"go"}
	r.POST("/form", handleForm)

	return r.Run(":8083")
}
