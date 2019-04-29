package api

import (
	"go_demo1/src/server/controller"
	"gopkg.in/kataras/iris.v5"
)

func Api() {

	var (
		demo = &controller.UserController{}
	)
	api := iris.Party("/demo1")
	api.Post("/hello", demo.HelloWorld)
}
