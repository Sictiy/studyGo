package _interface

type IAction interface {
	Init() bool
	DoAction() bool
	GetActionState() uint32
	SetPlayer(player *IPlayer)
	GetStateChan() chan uint32
	SetActionState(state uint32)
	GetTimeOut() uint32
	Config(attributies []string) bool
	GetIndex() uint8
	SetIndex(index uint8)
	SetActionName(string)
	GetActionName() string
	SetActionId(uint32)
	GetActionId() uint32
}


