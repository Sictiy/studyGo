package factory

import (
	"awesomeProject/src/entry/impl"
	. "awesomeProject/src/entry/impl/action"
	"awesomeProject/src/entry/interface"
	"awesomeProject/src/util"
	"strconv"
)

var roleIndex = 0

func NewPlayer() *_interface.IPlayer {
	pPlayer := &impl.Player{}
	pPlayer.SetAccountName("robot_"+strconv.Itoa(roleIndex))
	iPlayer := _interface.IPlayer(pPlayer)
	roleIndex++
	return &iPlayer
}

func NewActionByName(actionName string) *_interface.IAction {
	pIAction := NewAction(ActionId[actionName])
	if pIAction == nil {
		util.LogFormat("new action failed id: %d, name: %s", ActionId[actionName], actionName)
	}
	(*pIAction).SetActionName(actionName)
	(*pIAction).SetActionId(ActionId[actionName])
	return pIAction
}

func NewAction(actionName uint32) *_interface.IAction {
	var iAction _interface.IAction
	switch actionName {
	case ActionId["getRegionListReq"]:
		iAction = &GetRegionListReqAction{}
	case ActionId["loginReq"]:
		iAction = &LoginReqAction{}
	case ActionId["connectToServer"]:
		iAction = &ConnectToServer{}
	case ActionId["enterGameReq"]:
		iAction = &EnterGameReq{}
	}
	return &iAction
}

func NewRule(curAction uint32, curFinishCode uint32, nextAction uint32) *_interface.IRule {
	pRule := &impl.Rule{}
	pRule.SetCurAction(curAction)
	pRule.SetFinishCode(curFinishCode)
	pRule.SetNextAction(nextAction)
	iRule := _interface.IRule(pRule)
	return &iRule
}
