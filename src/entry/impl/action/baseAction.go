package action

import (
	mInterface "awesomeProject/src/entry/interface"
	"awesomeProject/src/util"
)

type BaseAction struct {
	pPlayer *mInterface.IPlayer
	actionState uint32
	stateChan chan uint32
	timeOut uint32 // 超时时间，秒
	index uint8
	actionName string
	actionId uint32
}

func (bAction *BaseAction) SetPlayer(player *mInterface.IPlayer) {
	bAction.pPlayer = player
}

func (bAction *BaseAction) GetPlayer() mInterface.IPlayer {
	return *bAction.pPlayer
}

func (bAction *BaseAction) SendMsg() {
	util.LogFormat("%d sendMsg",bAction.GetActionName())
}

func (bAction *BaseAction) SetActionName(name string)  {
	bAction.actionName = name
}

func (bAction *BaseAction) GetActionName() string {
	return  bAction.actionName
}

func (bAction *BaseAction) Init() bool {
	util.LogFormat("%d init", bAction.GetActionName())
	return true
}

func (bAction *BaseAction) DoAction() bool {
	util.LogFormat("%d doAction", bAction.GetActionName())
	return true
}

func (bAction *BaseAction) GetActionState() uint32 {
	return 0
}

func (bAction *BaseAction) GetStateChan() chan uint32 {
	return bAction.stateChan
}

func (bAction *BaseAction) SetActionState(state uint32) {
	bAction.actionState = state
}

func (bAction *BaseAction) GetTimeOut() uint32 {
	return bAction.timeOut
}

func (bAction *BaseAction) Config(attributies []string) bool {
	return true
}

func (bAction *BaseAction) GetIndex() uint8{
	return bAction.index
}

func (bAction *BaseAction) SetIndex(index uint8) {
	bAction.index = index
}

func (bAction *BaseAction) SetActionId(id uint32) {
	bAction.actionId = id
}

func (bAction *BaseAction) GetActionId() uint32 {
	return bAction.actionId
}

