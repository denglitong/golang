package examples

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type StructA struct {
	FieldA string `form:"field_a"`
}

type StructB struct {
	NestedStruct StructA
	FieldB       string `form:"field_b"`
}

type StructC struct {
	NestedStructPointer *StructA
	FieldC              string `form:"field_c"`
}

type StructD struct {
	NestedAnonyStruct struct {
		FieldX       string `form:"field_x"`
		NestedStruct StructA
	}
	FieldD string `form:"field_d"`
}

func getDataB(context *gin.Context) {
	var b StructB
	err := context.Bind(&b)
	if err != nil {
		log.Fatalln(err)
	}
	context.JSON(http.StatusOK, gin.H{
		"a": b.NestedStruct,
		"b": b.FieldB,
	})
}

func getDataC(context *gin.Context) {
	var c StructC
	err := context.Bind(&c)
	if err != nil {
		log.Fatalln(err)
	}
	context.JSON(http.StatusOK, gin.H{
		"a": c.NestedStructPointer,
		"c": c.FieldC,
	})
}

func getDataD(context *gin.Context) {
	var d StructD
	err := context.Bind(&d)
	if err != nil {
		log.Fatalln(err)
	}
	context.JSON(http.StatusOK, gin.H{
		"x": d.NestedAnonyStruct,
		"d": d.FieldD,
	})
}

func ShowBindDataRequestWithCustomStruct() error {
	r := gin.Default()

	// curl "http://localhost:8082/getb?field_a=hello&field_b=world"
	// return: {"a":{"FieldA":"hello"},"b":"world"}
	r.GET("/getb", getDataB)

	// curl "http://localhost:8082/getc?field_a=hello&field_c=world"
	// return: {"a":{"FieldA":"hello"},"c":"world"}
	r.GET("/getc", getDataC)

	// curl "http://localhost:8082/getd?field_x=hello&field_d=world&field_a=go"
	// return: {"d":"world","x":{"FieldX":"hello","NestedStruct":{"FieldA":"go"}}}
	r.GET("/getd", getDataD)

	return r.Run(":8082")
}
