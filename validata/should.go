package validata

import (
	"github.com/gin-gonic/gin"
	"log"
)

func ShouldBin(c *gin.Context, data interface{}) {
	err := c.ShouldBind(data)
	if err != nil {
		log.Fatalf("shouldBin parse data error=%v", err)
		return
	}
}
