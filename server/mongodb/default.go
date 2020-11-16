package mongodbServer

import (
	"gopkg.in/mgo.v2"
	"gin/config"
	"github.com/spf13/viper"
)


var session *mgo.Session
var database *mgo.Database

func init() {
	var err error
	if err = config.Init(""); err != nil {
		panic(err)
	}


	addr := viper.GetString("mongo.uri")
	user := viper.GetString("mongo.user")
	pwd := viper.GetString("mongo.pwd")
	databases := viper.GetString("mongo.database")
	if	user != "" && pwd != "" {
		session, err = mgo.Dial("mongodb://"+ user +":" + pwd + "@"+ addr)
	} else {
		session, err = mgo.Dial("mongodb://"+ addr)
	}

	database = session.DB(databases)
}

func GetMgo() *mgo.Session {
	return session
}

func GetDataBase() *mgo.Database {
	return database
}

func GetCollection(name string) *mgo.Collection {
	return database.C(name)
}

func GetErrNotFound() error {
	return mgo.ErrNotFound
}
