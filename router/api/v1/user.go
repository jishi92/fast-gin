package v1

import (
	"fast-gin/model"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"net/http"

	"fast-gin/ecode"
	"fast-gin/router/api"
	"fast-gin/service"
	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {
	var (
		req model.AddUserReq
		err error
	)

	//httpCode, errCode := api.BindAndValid(c, &req)
	err = c.BindWith(&req, binding.FormPost)
	if err != nil {
		fmt.Println("err===", err)
		return
	}
	fmt.Println("req===", req)

	err = service.AddUser(req)
	if err != nil {
		api.Response(c, http.StatusInternalServerError, ecode.AddUserFail, nil)
		return
	}

	api.Response(c, http.StatusOK, ecode.SUCCESS, nil)
}

func UpdateUser(c *gin.Context) {
	var (
		req model.UpdateUserReq
		err error
	)
	httpCode, errCode := api.BindAndValid(c, &req)
	if errCode != ecode.SUCCESS {
		api.Response(c, httpCode, errCode, nil)
		return
	}

	err = service.UpdateUser(req)
	if err != nil {
		api.Response(c, http.StatusInternalServerError, ecode.EditUserFail, nil)
		return
	}

	api.Response(c, http.StatusOK, ecode.SUCCESS, nil)
}

func GetUserInfo(c *gin.Context) {
	var (
		req model.IdReq
		err error
	)
	httpCode, errCode := api.BindAndValid(c, &req)
	if errCode != ecode.SUCCESS {
		api.Response(c, httpCode, errCode, nil)
		return
	}

	reply, err := service.GetUserInfo(req.ID)
	if err != nil {
		api.Response(c, http.StatusInternalServerError, ecode.GetUserInfoFail, nil)
		return
	}

	api.Response(c, http.StatusOK, ecode.SUCCESS, reply)
}

func GetAllUsers(c *gin.Context) {
	var (
		req model.UserListReq
		err error
	)
	httpCode, errCode := api.BindAndValid(c, &req)
	if errCode != ecode.SUCCESS {
		api.Response(c, httpCode, errCode, nil)
		return
	}

	users, err := service.GetAllUsers(req)
	if err != nil {
		api.Response(c, http.StatusInternalServerError, ecode.GetAllUserFail, nil)
		return
	}

	api.Response(c, http.StatusOK, ecode.SUCCESS, users)
}

func DelUser(c *gin.Context) {
	var (
		req model.IdReq
		err error
	)
	httpCode, errCode := api.BindAndValid(c, &req)
	if errCode != ecode.SUCCESS {
		api.Response(c, httpCode, errCode, nil)
		return
	}

	err = service.DelUser(req.ID)
	if err != nil {
		api.Response(c, http.StatusInternalServerError, ecode.DeleteUserFail, nil)
		return
	}

	api.Response(c, http.StatusOK, ecode.SUCCESS, nil)
}
