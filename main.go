package main

import (
	"gin/config"
	"gin/lib/logging"
	"gin/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"net/http"
)


func main() {
	pflag.Parse()
	if err := config.Init(""); err != nil {
		panic(err)
	}
	gin.SetMode(viper.GetString("runmode"))

	g := gin.New()
	g.Use(logging.Logger())
	g.Use(gin.Recovery())
	router.Load(g)
	_ = http.ListenAndServe(viper.GetString("addr"), g)
}