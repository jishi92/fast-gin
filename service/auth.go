package service

import (
	"fast-gin/dao"
)

type Auth struct {
	Username string
	Password string
}

func (a *Auth) Check() (bool, error) {
	return dao.CheckAuth(a.Username, a.Password)
}
