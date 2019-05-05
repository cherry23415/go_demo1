package dao

import (
	"encoding/json"
	"github.com/bitly/go-simplejson"
	"github.com/go-errors/errors"
	"go_demo1/src/server/common"
)

type UserDao struct {
}

type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Age      int64  `json:"age"`
	Birthday string `json:"birthday"`
}

func (*UserDao) AddUser(data []byte) {
	var user User
	json.Unmarshal(data, &user)
	common.DBGorm.Table("t_user").Create(user)
	return
}

func (*UserDao) UpdateUser(data []byte) error {
	var user User
	json.Unmarshal(data, &user)
	if user.Id <= 0 {
		return errors.New("ID错误")
	}
	common.DBGorm.Table("t_user").Model(&user).Updates(user)
	return nil
}

func (*UserDao) AllUser() (*simplejson.Json, error) {
	var users []User
	common.DBGorm.Table("t_user").Find(&users)
	jsonBytes, err2 := json.Marshal(users)
	if err2 != nil {
		return nil, err2
	}
	return simplejson.NewJson(jsonBytes)
}

func (*UserDao) GetUserById(id int64) (*simplejson.Json, error) {
	var user User
	common.DBGorm.Table("t_user").Find(&user, id)
	jsonBytes, err2 := json.Marshal(user)
	if err2 != nil {
		return nil, err2
	}
	return simplejson.NewJson(jsonBytes)
}
