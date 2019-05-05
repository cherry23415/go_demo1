package service

import (
	"go_demo1/src/server/dao"
)

type UserService struct {
}

func (*UserService) AddUser(data []byte) {
	userDao := dao.UserDao{}
	userDao.AddUser(data)
	return
}
