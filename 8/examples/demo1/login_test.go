package demo1

import "testing"

func TestLogin1(t *testing.T) {
	user := "root"
	pass := "123456"
	isLogin := Login(user, pass)
	// 登录失败
	if !isLogin {
		t.Errorf("user=%s,pass=%s, 用户名必须是root,密码必须是123456", user, pass)
	}
}

func TestLogin2(t *testing.T) {
	user := "admin"
	pass := "123456"
	isLogin := Login(user, pass)
	// 登录失败
	if !isLogin {
		t.Errorf("user=%s,pass=%s, 用户名必须是root,密码必须是123456", user, pass)
	}
}