package service

import (
	"fast-gin/dao"
	"fast-gin/model"
	"fmt"
	"log"
)

func AddUser(req model.AddUserReq) error {
	m := model.User{
		Name:   req.Name,
		Remark: req.Remark,
		State:  1,
	}
	return dao.AddUser(m)
}

func UpdateUser(req model.UpdateUserReq) error {
	data := map[string]interface{}{
		"username": req.Name,
		"remark":   req.Remark,
	}
	if req.State != 0 {
		data["state"] = req.State
	}
	return dao.UpdateUser(req.Id, data)
}

func GetUserInfo(id int64) (reply model.UserInfoReply, err error) {
	user, err := dao.GetUser(id)
	if err != nil || user == nil {
		log.Printf("GetUserInfo req(%v) user(%v) err(%v)", id, user, err)
		return
	}
	reply.Id = user.ID
	reply.Name = user.Name
	reply.Remark = user.Remark
	reply.State = user.State
	return
}

func DelUser(id int64) (err error) {
	return dao.DeleteUser(id)
}

func GetAllUsers(req model.UserListReq) (reply *model.UserListReply, err error) {
	m := make(map[string]interface{})
	reply = new(model.UserListReply)
	if req.State != 0 {
		m["state"] = req.State
	}
	if req.Name != "" {
		m["username"] = req.Name
	}
	total, err := dao.CountUsers(m)
	if err != nil || total == 0 {
		return
	}
	reply.Total = total
	reply.Pn = req.Pn
	reply.Ps = req.Ps
	users, err := dao.GetAllUser(m, req.Offset(), req.Ps)
	fmt.Println("get=====list", req, m, users)
	if err != nil {
		fmt.Println("GetAllUser err", err)
		return
	}
	var list []*model.UserInfoReply
	for _, v := range users {
		tmp := &model.UserInfoReply{
			Id:     v.ID,
			Name:   v.Name,
			Remark: v.Remark,
			State:  v.State,
		}
		list = append(list, tmp)
	}
	reply.List = list
	return
}
