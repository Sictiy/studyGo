package main

import "awesomeProject/src/util"

func main() {
	type mStruct struct {
		mInt uint32
	}
	mMap := make(map[uint32]*mStruct)
	mMap[1] = &mStruct{1}
	util.Log(mMap[1])
	util.Log(mMap[2])
	if mMap[2] == nil {
		util.LogFormat("空")
	}
	util.Log(mMap)
}
