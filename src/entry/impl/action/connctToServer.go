package action

import (
	"awesomeProject/src/net"
	"awesomeProject/src/util"
	"fmt"
)

type ConnectToServer struct {
	BaseAction
	ip string
	port string
	network string
}

func (bAction *ConnectToServer) Init() bool {
	bAction.network = "tcp"
	if bAction.GetPlayer().GetCurSession() == nil {
		bAction.ip = "128.1.43.103"
		bAction.port = "20001"
	}else{
		bAction.ip = "128.1.43.103"
		bAction.port = "20002"
	}
	util.LogFormat("init connect to server action ip:%s  port: %s", bAction.ip, bAction.port)
	return true
}

func (bAction *ConnectToServer) DoAction() bool {
	session := net.ConnetTo(bAction.network, bAction.ip+":"+bAction.port)
	if session == nil {
		util.LogFormat("连接到服务器失败")
		return false
	}
	util.LogFormat("%s connect to server: %s", bAction.GetPlayer().GetAccountName(), fmt.Sprintf("%s:%s", bAction.ip, bAction.port))
	bAction.GetPlayer().SetCurSession(session)
	return true
}
