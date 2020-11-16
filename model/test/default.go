package test

import (
	"gin/server/mongodb"
	"gin/server/mysql"
	"gin/server/mysql/schema"
	"time"
)
func MysqlTest(){
	db := mysqlServer.GetDatabase()
	mogodb := mongodbServer.GetDataBase()
	mogodb.C("test")
	//db.Model(&schema.CreditCard{})
	db.Create(&schema.User{
		UserID: uint(time.Now().Second()),
		Number: "kdjkfd",
	})
}
