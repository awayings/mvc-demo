package main

import "mvc-demo/pkg/service"

func main() {

	//service.MsgService.Persist(nil, nil)
	//
	//service.MsgSvc.Persist(nil, nil)

	service.MsgModelRWIface.Persist(nil, nil)
	service.MsgModelIface.Persist(nil, nil)
}