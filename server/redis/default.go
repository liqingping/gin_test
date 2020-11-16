package redisServer

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gin/config"
)

var Client0 *redis.Client
var Client1 *redis.Client
var Client2 *redis.Client

func init(){
	if err := config.Init(""); err != nil {
		panic(err)
	}
	addr :=viper.GetString("redis.host")+":"+viper.GetString("redis.port")
	pwd:=viper.GetString("redis.pwd")

	if pwd != "" {
		Client0 = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: pwd, // no password set
			DB:       0,  // use default DB
		})
		Client1 = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: pwd, // no password set
			DB:       1,  // use default DB
		})
		Client2 = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: pwd, // no password set
			DB:       2,  // use default DB
		})
	} else {
		Client0 = redis.NewClient(&redis.Options{
			Addr:     addr,
			DB:       0,  // use default DB
		})
		Client1 = redis.NewClient(&redis.Options{
			Addr:     addr,
			DB:       1,  // use default DB
		})
		Client2 = redis.NewClient(&redis.Options{
			Addr:     addr,
			DB:       2,  // use default DB
		})
	}

}

func GetClient(num int) *redis.Client{
	var client *redis.Client
	if num == 0 {
		client= Client0
	}
	if num == 1 {
		client= Client1
	}
	if num == 2 {
		client= Client2
	}

	return client
}