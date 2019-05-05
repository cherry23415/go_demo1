package dao

import (
	"encoding/json"
	"go_demo1/src/server/common"
)

type UserDao struct {
}

type User struct {
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
