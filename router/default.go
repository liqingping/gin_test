package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gin/router/common"
	"gin/router/test"
)

//处理跨域问题
func Cross() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		//设置放回的header头
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Authorization, Accept, X-Requested-With , yourHeaderFeild")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Allow-Credentials", "true")
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
		} else {
			c.Next()
		}
	}
}

func Load(g *gin.Engine) *gin.Engine {
	g.Use(Cross())
	common.MapRoute(g)
	test.MapRoute(g)
	return g
}
