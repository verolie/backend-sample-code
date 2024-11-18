package product

import "github.com/gin-gonic/gin"

func PatchProduct(c *gin.Context) {
	var b StructC
	c.Bind(&b)

	c.JSON(200, gin.H{
		"a": b.NestedStructPointer,
		"c": b.FieldC,
	})
}