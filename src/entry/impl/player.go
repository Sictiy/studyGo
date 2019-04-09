package impl

import (
	"awesomeProject/src/Proto"
	"awesomeProject/src/entry/interface"
	"awesomeProject/src/mgr"
	"awesomeProject/src/net"
	"awesomeProject/src/util"
	"time"
	)

type Player struct {
	accountId uint64
	accountName string
	roleId uint64
	roleName string
	job  int32
	sex  int32
	curSession *net.RbSession // 当前连接的网络session
	curAction *_interface.IAction //当前执行的action
	lastAction *_interface.IAction // 上次执行的action
}

func (pPlayer *Player) SetAccountName(name string) {
	pPlayer.accountName = name
}

func (pPlayer *Player) SetRoleName(name string) {
	pPlayer.roleName = name
}

func (pPlayer *Player) SetAccountId(id uint64) {
	pPlayer.accountId = id
}

func (pPlayer *Player) SetRoleId(id uint64) {
	pPlayer.roleId = id
}

func (pPlayer *Player) GetAccountName() string {
	return pPlayer.accountName
}

func (pPlayer *Player) GetRoleName() string {
	return pPlayer.roleName
}

func (pPlayer *Player) GetAccountId() uint64 {
	return pPlayer.accountId
}

func (pPlayer *Player) GetRoleId() uint64 {
	return pPlayer.roleId
}

func (pPlayer *Player) GetCurSession() *net.RbSession {
	return pPlayer.curSession
}

func (pPlayer *Player) SetCurSession(pSession *net.RbSession) {
	if pPlayer.curSession != nil {
		pPlayer.curSession.Disconnect()
	}
	pPlayer.curSession = pSession
}

func (pPlayer *Player) ReDoAction(){
	pPlayer.curAction = nil
	pPlayer.StartDoAction()
}

func (pPlayer *Player) StartDoAction() {
	if pPlayer.curAction != nil {
		return
	}
	util.LogFormat("%s start do action", pPlayer.accountName)
	firstAction := mgr.GetActionsMgr().GetFirstAction()
	pPlayer.SetCurActionAndDo(firstAction)
}

func (pPlayer *Player) DoNextAction() {
	pPlayer.lastAction = pPlayer.curAction
	pPlayer.curAction = nil
	nextAction := mgr.GetActionsMgr().GetNetAction(*pPlayer.lastAction)
	if nextAction != nil {
		util.LogFormat("%s do nextAction", pPlayer.accountName)
		pPlayer.SetCurActionAndDo(nextAction)
	} else{
		util.LogFormat("%s do action finished!", pPlayer.GetAccountName())
	}
}

func (pPlayer *Player) SetCurActionAndDo(pAction *_interface.IAction) {
	var curAction = *pAction
	iPlayer := _interface.IPlayer(pPlayer)
	curAction.SetPlayer(&iPlayer)
	pPlayer.curAction = pAction
	util.LogFormat("setCurAction actionName: %s", curAction.GetActionName())
	curAction.Init()
	curAction.DoAction()
	go pPlayer.checkActionState()
}

func (pPlayer *Player) checkActionState() {
	// select 阻塞协程
	curAction := *pPlayer.curAction
	stateChan := curAction.GetStateChan()
	var state uint32
	select {
	case state = <- stateChan:
		util.LogFormat("state: %d", state)
	case <- time.After(time.Second * 40):
		util.LogFormat("timeout  state: %d")
		state = Proto.ACTION_STATE_TIME_OUT
	}
	curAction.SetActionState(state)
	pPlayer.DoNextAction()
}

func (pPlayer *Player) SendMsg(cmd uint32, data []byte) {
	pack := net.NewRbPackage()
	pack.AddData(data)
	pack.SetCmd(cmd)
	pPlayer.GetCurSession().SendPackage(pack)
}

func (pPlayer *Player) SendPackage(rbp net.RbPackage) {
	pPlayer.GetCurSession().SendPackage(rbp)
}
