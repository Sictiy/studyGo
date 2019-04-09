package net

func ConnetTo(network string, ipport string) *RbSession {
	var rbSession RbSession
	ret := rbSession.ConnectToServer(network, ipport)
	if !ret {
		return nil
	}
	rbSession.StartRecv()
	return &rbSession
}
