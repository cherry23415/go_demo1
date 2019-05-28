package initialize

import (
	"github.com/spf13/viper"
	"go_gorm_demo/src/server/route/api"
	"gopkg.in/kataras/iris.v5"
)

func SetupServer() {

	port := viper.GetString("server.port")
	host := viper.GetString("server.host")

	api.Api()

	iris.Listen(host + ":" + port)

}
