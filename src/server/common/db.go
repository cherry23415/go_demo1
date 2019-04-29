package common

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

func SetUpDBGorm() {

	dialect := viper.GetString("server.db.dialect")
	user := viper.GetString("server.db.user")
	password := viper.GetString("server.db.password")
	database := viper.GetString("server.db.database")
	host := viper.GetString("server.db.host")
	port := viper.GetString("server.db.port")
	maxIdle := viper.GetInt("server.db.maxIdle")
	maxOpen := viper.GetInt("server.db.maxOpen")

	url := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(dialect, url)
	if err != nil {
		panic("failed to connect database")
	}
	db.DB().SetMaxOpenConns(maxOpen) //最大连接数
	db.DB().SetMaxIdleConns(maxIdle) //空闲连接数
	connectErr := db.DB().Ping()     //测试连接
	if connectErr != nil {
		panic("failed to connect database:" + connectErr.Error())
	}
	DBGorm = db
}

var (
	DBGorm *gorm.DB
)
