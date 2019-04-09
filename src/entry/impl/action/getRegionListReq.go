package action

import (
	"awesomeProject/src/Proto"
	"awesomeProject/src/util"
	"github.com/google/flatbuffers/go"
)

type GetRegionListReqAction struct {
	BaseAction
}

func (bAction *GetRegionListReqAction) DoAction() bool {
	iPlayer := bAction.GetPlayer()
	build := flatbuffers.NewBuilder(0)
	str := build.CreateString(iPlayer.GetAccountName())
	Proto.GetRegionListReqStart(build)
	Proto.GetRegionListReqAddAccount(build, str)
	strSend := Proto.GetRegionListReqEnd(build)
	build.Finish(strSend)

	iPlayer.SendMsg(Proto.CMDC_LS_GET_REGION_LIST_REQ, build.FinishedBytes())
	util.LogFormat("%s send get region list req")
	return true
}
