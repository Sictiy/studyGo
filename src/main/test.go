package main

import (
	"awesomeProject/src/util"
	"fmt"
	"time"
)

func main() {
	test2()
}

func test1()  {
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

func test2()  {
	var m uint64
	var n uint64
	fmt.Scanln(&m)
	fmt.Scanln(&n)
	util.LogFormat("m: %d, n: %d", m, n)
	resultChan := make(chan uint64)
	go func(){resultChan <- A(m ,n)}()
	startTime := time.Now()
	Loop:
	for {
		select {
		case result := <- resultChan :
			util.LogFormat("A(%d,%d) = %d", m, n, result)
			break Loop
		case <- time.After(time.Second * 2):
			util.LogFormat("runing waited times: %ds", time.Now().Sub(startTime) / time.Second)
		}
	}
}

func A(m uint64, n uint64) uint64 {
	for m != 0 {
		if n == 0 {
			n = 1
		} else{
			n = A(m, n-1)
		}
		m -= 1
	}
	return n+1
}