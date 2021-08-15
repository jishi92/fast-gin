package router

import (
	"fmt"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	e "fast-gin/ecode"
	"fast-gin/library/util"
	"fast-gin/router/api"
	"fast-gin/service"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// @Summary Get Auth
// @Produce  json
// @Param username query string true "userName"
// @Param password query string true "password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth [get]
func GetAuth(c *gin.Context) {
	valid := validation.Validation{}

	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println("auth====", valid, username, password)
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	if !ok {
		api.Response(c, http.StatusBadRequest, e.InvalidParams, nil)
		return
	}

	authService := service.Auth{Username: username, Password: password}
	isExist, err := authService.Check()
	if err != nil {
		api.Response(c, http.StatusInternalServerError, e.ErrorAuthCheckTokenFail, nil)
		return
	}

	if !isExist {
		api.Response(c, http.StatusUnauthorized, e.ErrorAuth, nil)
		return
	}

	token, err := util.GenerateToken(username, password)
	if err != nil {
		api.Response(c, http.StatusInternalServerError, e.ErrorAuthToken, nil)
		return
	}

	api.Response(c, http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})
}
