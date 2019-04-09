package _interface

type IRule interface {
	GetNextActionId() uint32
	GetCurActionId() uint32
	GetFinishCode() uint32
}
