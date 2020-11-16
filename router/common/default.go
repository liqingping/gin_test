package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func MapRoute(g *gin.Engine) {
	g.HEAD("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": "0000",
		})
	})

	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})
}
