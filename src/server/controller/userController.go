package controller

import (
	"go_demo1/src/server/errcode"
	"gopkg.in/kataras/iris.v5"
)

type UserController struct {
}

func (*UserController) HelloWorld(ctx *iris.Context) {
	ctx.JSON(iris.StatusOK, errcode.SUCCESS.Result())
	return

}
