package main

type RobotTool struct {
	stopChan chan uint32
}

func (rt *RobotTool) StartRun()  {
	rt.stopChan = make(chan uint32)
}

func (rt *RobotTool) StopRun()  {
	rt.stopChan <- 0
}

func (rt *RobotTool) GetStopChan() chan uint32{
	return rt.stopChan
}
