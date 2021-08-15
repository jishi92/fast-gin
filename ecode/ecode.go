package ecode

// 自定义的错误码
const (
	SUCCESS       = 200
	InvalidParams = 400
	Error         = 500

	GetUserInfoFail = 10001
	GetAllUserFail  = 10002
	AddUserFail     = 10003
	EditUserFail    = 10004
	DeleteUserFail  = 10005

	ErrorAuthCheckTokenFail    = 20001
	ErrorAuthCheckTokenTimeout = 20002
	ErrorAuthToken             = 20003
	ErrorAuth                  = 20004

	ErrorUploadSaveImageFail    = 30001
	ErrorUploadCheckImageFail   = 30002
	ErrorUploadCheckImageFormat = 30003
)

// 错误码对应的错误消息

var Msg = map[int]string{
	SUCCESS:       "成功",
	Error:         "错误",
	InvalidParams: "请求参数错误",

	GetUserInfoFail: "获取用户信息失败",
	GetAllUserFail:  "获取用户列表失败",
	AddUserFail:     "新增用户失败",
	EditUserFail:    "更新用户失败",
	DeleteUserFail:  "删除用户失败",

	ErrorAuthCheckTokenFail:    "Token鉴权失败",
	ErrorAuthCheckTokenTimeout: "Token已超时",
	ErrorAuthToken:             "Token生成失败",
	ErrorAuth:                  "Token错误",

	ErrorUploadSaveImageFail:    "保存图片失败",
	ErrorUploadCheckImageFail:   "检查图片失败",
	ErrorUploadCheckImageFormat: "校验图片错误，图片格式或大小有问题",
}

func GetMsg(code int) string {
	msg, ok := Msg[code]
	if ok {
		return msg
	}

	return Msg[Error]
}
