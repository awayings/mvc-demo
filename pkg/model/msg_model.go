package model

import "context"

// 定义一个基础model
type MsgModel interface {
	MsgModelRW
	GetList(context context.Context, uid, peerId, sinceMsgId, maxMsgId int64, count int) (interface{}, bool)
}

type MsgModelRW interface {
	Persist(context context.Context, msg interface{}) bool
	UpdateDbContent(context context.Context, msgIface interface{}) bool
}
