package net

import (
	"awesomeProject/src/util"
	"io/ioutil"
	"net"
)

type RbSession struct {
	pTcpAddr *net.TCPAddr
	pTcpConn *net.TCPConn
}

func (rbSession *RbSession) ConnectToServer(network string, ipport string) bool {
	var err error
	rbSession.pTcpAddr, err = net.ResolveTCPAddr(network, ipport)
	util.CheckError(err)
	rbSession.pTcpConn, err = net.DialTCP(network, nil, rbSession.pTcpAddr)
	util.CheckError(err)
	return true
}

func (rbSession *RbSession) StartRecv() {
	go rbSession.loopRecv()
}

func (RbSession) Disconnect() bool{
	return true
}

func (rbSession *RbSession) SendPackage(rbPackage RbPackage) int32 {
	return rbSession.SendBytes(rbPackage.GetData())
}

func (rbSession *RbSession) SendBytes(data []byte) int32 {
	return 0
}

func (rbSession *RbSession) loopRecv() {
	for {
		data, err := ioutil.ReadAll(rbSession.pTcpConn)
		util.LogFormat("recvData： %s", data)
		util.CheckError(err)
		pack := NewRbPackageFormBytes(data)
		if len(pack.GetData()) <= 0 {
			break
		}
		rbSession.onRecv(pack)
	}
}

func (RbSession) onRecv(rbPackage RbPackage){
	
}

func (RbSession) onDisconnect() {

}

func (RbSession) onConnect() {

}