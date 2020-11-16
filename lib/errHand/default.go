package errHand

import (
	"github.com/gin-gonic/gin"
	"gin/lib/logging"
)

func Panic(c *gin.Context)  {
	if err := recover(); err != nil {
		res := make(map[string]interface{})
		res["code"] = "0100"
		res["msg"] = err
		if err, ok := err.(error); ok {
			res["msg"] = err.Error()
		}
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		if raw != "" {
			path = path + "?" + raw
		}
		logging.Panic(err, path)
		c.JSON(400, res)
	}
}

func Validator(err error) {
	if err != nil {
		panic(err)
	}
}