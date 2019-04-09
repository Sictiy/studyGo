package _interface

import "awesomeProject/src/net"

type IPlayer interface {
	GetAccountName() string
	GetRoleName() string
	GetAccountId() uint64
	GetRoleId() uint64
	GetCurSession() *net.RbSession
	SetCurSession(*net.RbSession)

	StartDoAction()
	DoNextAction()
	SetCurActionAndDo(action *IAction)

	SendMsg(uint32, []byte)
	SendPackage(net.RbPackage)
}
