package mysqlServer

import (
	"gin/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var DB *gorm.DB
func init() {
	if err := config.Init(""); err != nil {
		panic(err)
	}
	uri :=viper.GetString("mysql.user")+":"+
		viper.GetString("mysql.pwd")+"@tcp("+
		viper.GetString("mysql.host")+":"+
		viper.GetString("mysql.port")+")/"+
		viper.GetString("mysql.database")+"?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open("mysql", uri)
	if err != nil {
		db.Close()
		return
	}
	// 启用Logger，显示详细日志
	db.LogMode(true)
	// 全局禁用表名复数
	db.SingularTable(true)

	//db.AutoMigrate(&schema.User{})

	DB = db
}


func GetDatabase() *gorm.DB {
	return DB
}
