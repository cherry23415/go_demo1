package service

import (
	"github.com/bitly/go-simplejson"
	"go_demo1/src/server/dao"
)

type UserService struct {
}

func (*UserService) AddUser(data []byte) {
	userDao := dao.UserDao{}
	userDao.AddUser(data)
	return
}

func (*UserService) UpdateUser(data []byte) error {
	userDao := dao.UserDao{}
	return userDao.UpdateUser(data)
}

func (*UserService) AllUser() (*simplejson.Json, error) {
	userDao := dao.UserDao{}
	return userDao.AllUser()
}

func (*UserService) GetUserById(id int64) (*simplejson.Json, error) {
	userDao := dao.UserDao{}
	return userDao.GetUserById(id)
}
