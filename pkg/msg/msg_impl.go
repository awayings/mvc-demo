package msg

import "context"

type msgModelImpl struct{}
var MsgModelImpl = msgModelImpl{}
func (m *msgModelImpl) Persist(context context.Context, msgIface interface{}) bool {
	// 具体实现
	println("persist")
	return true
}

func (m *msgModelImpl) UpdateDbContent(context context.Context, msgIface interface{}) bool {
	// 具体实现
	println("update")
	return true
}

func (m *msgModelImpl) GetList(context context.Context, uid, peerId, sinceMsgId, maxMsgId int64, count int) (interface{}, bool){
	// 具体实现
	println("context")
	return nil, true
}

