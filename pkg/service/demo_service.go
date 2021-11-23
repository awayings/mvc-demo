package service

import (
	"mvc-demo/pkg/model"
	"mvc-demo/pkg/msg"
)

// 定义一个msgService struct包含了model里面的UserModel和MsgModel两个model
type msgService struct {
//	userModel  model.UserModel
	msgModel   model.MsgModel
}

type msgSvc struct {
	model.MsgModel
}

var MsgSvc = msgSvc{&msg.MsgModelImpl}

// 定义一个MsgService的变量，并初始化，这样通过MsgService，就能引用并访问model的所有方法
var (
	MsgService = msgService{
		//userModel:      user.UserModelImpl,
		msgModel:       &msg.MsgModelImpl,
	}
)

// 以下说明不同接口之间赋值的情况。
var MsgModelIface model.MsgModel = &msg.MsgModelImpl
var MsgModelRWIface model.MsgModelRW = MsgModelIface
