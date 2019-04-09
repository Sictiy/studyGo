package main

import (
	"awesomeProject/src/Proto"
	"awesomeProject/src/entry/factory"
	"awesomeProject/src/mgr"
	"awesomeProject/src/util"
	"time"
)

var stopChan chan uint32

func main() {
	//StartRun()
	stopChan = make(chan uint32)
	defer close(stopChan)

	util.LogFormat("开始测试 go")
	pActionsMgr := mgr.GetActionsMgr()
	pPlayerMgr := mgr.GetPlayerMgr()
	pActionsMgr.AddAction(factory.NewActionByName("connectToServer"))
	pActionsMgr.AddAction(factory.NewActionByName("getRegionListReq"))
	pActionsMgr.AddAction(factory.NewActionByName("loginReq"))
	pActionsMgr.AddAction(factory.NewActionByName("connectToServer"))
	pActionsMgr.AddAction(factory.NewActionByName("enterGameReq"))

	pActionsMgr.AddRule(factory.NewRule(0, Proto.ACTION_STATE_TIME_OUT, 0))
	for i := 0; i < 10 ; i++ {
		pPlayerMgr.AddPlayer(factory.NewPlayer())
	}
	pPlayerMgr.StartDoAction()

	go waitStop()

	select {
	case <- stopChan:
		util.LogFormat("测试结束 go")
	}
	//net.TestNet()
}

func waitStop() {
	time.Sleep(time.Second * 10)
	stopChan <- 10
}