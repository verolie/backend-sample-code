package product

import (
	"github.com/gin-gonic/gin"
)

type StructA struct {
    FieldA string `form:"field_a"`
}

type StructC struct {
    NestedStructPointer *StructA
    FieldC string `form:"field_c"`
}

func CreateProduct(c *gin.Context) {
	var b StructC
    c.Bind(&b)

    c.JSON(200, gin.H{
        "a": b.NestedStructPointer,
        "c": b.FieldC,
    })
}