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
	return
}

func (*UserController) UpdateUser(ctx *iris.Context) {
	putBody := ctx.PostBody()
	userService := service.UserService{}
	err := userService.UpdateUser(putBody)
	if err != nil {
		ctx.JSON(iris.StatusOK, errcode.PARAM_ERROR.ResultWithMsg(err.Error()))
		return
	}
	ctx.JSON(iris.StatusOK, errcode.SUCCESS.Result())
	return
}

func (*UserController) AllUser(ctx *iris.Context) {
	userService := service.UserService{}
	result, err := userService.AllUser()
	if err != nil {
		ctx.JSON(iris.StatusOK, errcode.PARAM_ERROR.ResultWithMsg(err.Error()))
		return
	}
	ctx.JSON(iris.StatusOK, errcode.SUCCESS.ResultWithData(result))
	return
}

func (*UserController) GetUserById(ctx *iris.Context) {
	id, err := ctx.URLParamInt64("id")
	if err != nil {
		ctx.JSON(iris.StatusOK, errcode.PARAM_ERROR.ResultWithMsg("参数错误"))
		return
	}
	userService := service.UserService{}
	result, err := userService.GetUserById(id)
	if err != nil {
		ctx.JSON(iris.StatusOK, errcode.PARAM_ERROR.ResultWithMsg(err.Error()))
		return
	}
	ctx.JSON(iris.StatusOK, errcode.SUCCESS.ResultWithData(result))
	return
}
