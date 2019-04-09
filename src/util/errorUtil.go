package util

import "fmt"

func CheckError(err error)  {
	if err != nil{
		fmt.Print(err)
	}
}

func LogFormat(format string, a ... interface{}){
	format = format + "\n"
	if a != nil{
		_,e := fmt.Printf(format, a ...)
		CheckError(e)
	}else{
		fmt.Print(format)
	}
}

func Log(a ... interface{})  {
	fmt.Print(a ...)
}
