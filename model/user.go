package model

import (
	"github.com/jinzhu/gorm"
)

/*
model模块用于记录结构体
包括表的结构体和请求结构体和返回结构体
*/

type User struct {
	gorm.Model
	Name   string `gorm:"column:username"`
	Pwd    string `gorm:"column:pwd"`
	Remark string `gorm:"column:remark"`
	State  int8   `gorm:"column:state"`
}

// 设置表名，可以通过给struct类型定义 TableName函数，返回当前struct绑定的mysql表名是什么

func (u User) TableName() string {
	return "user"
}

type PageReq struct {
	Pn int64 `json:"pn" form:"pn" default:"1"`
	Ps int64 `json:"ps" form:"ps" validate:"max=50" default:"20"`
}

func (p PageReq) Offset() (offset int64) {
	if p.Pn == 0 {
		offset = 0
	} else {
		offset = (p.Pn - 1) * p.Ps
	}
	return
}

type PageReply struct {
	Pn    int64 `json:"pn"`
	Ps    int64 `json:"ps"`
	Total int64 `json:"total"`
}

type AddUserReq struct {
	Name   string `form:"name" valid:"Required;MaxSize(100)"`
	Remark string `form:"remark"`
}

type UpdateUserReq struct {
	Id    int64 `form:"id"`
	State int8  `form:"state"`
	AddUserReq
}
type IdReq struct {
	ID int64 `form:"id"`
}

type UserListReq struct {
	Name  string `json:"username" form:"username"`
	State int    `json:"state" form:"state"`
	PageReq
}

type UserInfoReply struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Remark string `json:"remark"`
	State  int8   `json:"state"`
}

type UserListReply struct {
	List []*UserInfoReply
	PageReply
}
