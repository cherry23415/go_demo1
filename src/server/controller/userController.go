package controller

import (
	"go_demo1/src/server/errcode"
	"go_demo1/src/server/service"
	"gopkg.in/kataras/iris.v5"
)

type UserController struct {
}

func (*UserController) HelloWorld(ctx *iris.Context) {
	ctx.JSON(iris.StatusOK, errcode.SUCCESS.ResultWithData("Hello,this is go demo1!"))
	return
}

func (*UserController) AddUser(ctx *iris.Context) {
	putBody := ctx.PostBody()
	userService := service.UserService{}
	userService.AddUser(putBody)
	ctx.JSON(iris.StatusOK, errcode.SUCCESS.Result())
}
