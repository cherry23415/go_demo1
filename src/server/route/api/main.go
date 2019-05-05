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
	api.Get("/hello", demo.HelloWorld)
	api.Post("/addUser", demo.AddUser)
}
