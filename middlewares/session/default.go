package session

import (
	"encoding/json"
	"errors"
	"gin/server/redis"
	"github.com/gin-gonic/gin"
	"strings"
)

type Session struct {
	Id string `json:"id"`
	Token string `json:"token"`
	Expire int `json:"expire"`
}

func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				res := make(map[string]interface{})
				res["code"] = 0
				res["msg"] = err
				if err, ok := err.(error); ok {
					res["msg"] = err.Error()
				}
				c.JSON(404, res)
				c.Abort()
			}
		}()

		cookie := c.GetHeader("Authorization")
		if cookie == "" || cookie == "Bearer"{
			res := make(map[string]interface{})
			res["code"] = "0403"
			res["msg"] = "请登录"

			c.JSON(403, res)
			c.Abort()
			return
		}
		cookie = strings.Split(cookie, " ")[1]
		//cookie := "03529de30f9fa45cab5c96b2882b6a3a_U"

		client := redisServer.Client0
		val, err := client.Get(cookie).Result()

		if err != nil || val == "" {
			res := make(map[string]interface{})
			res["code"] = "0403"
			res["msg"] = "请登录"

			c.JSON(403, res)
			c.Abort()
			return
		} else {
			c.Set("session", val)
			c.Set("auth", true)
			c.Next();
		}
	}
}

func GetSession(c *gin.Context) (Session, error){
	auth := c.GetBool("auth")
	session := c.GetString("session")
	var s Session
	if !auth {
		cookie := c.GetHeader("Authorization")
		if cookie == "" || cookie == "Bearer"{
			return s, errors.New("未登录")
		}
		cookie = strings.Split(cookie, " ")[1]
		client := redisServer.Client0
		val, err := client.Get(cookie).Result()
		if err != nil || val == "" {
			return s, errors.New("未登录")
		} else {
			json.Unmarshal([]byte(val), &s)
			return s, nil
		}
	} else {
		json.Unmarshal([]byte(session), &s)
		return s, nil
	}
}
