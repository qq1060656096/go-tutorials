// examples/demo2/login_test.go
package demo2

import (
	"fmt"
	"testing"
)

// 成功的测试
func TestLoginSuccess(t *testing.T) {
	user := "root"
	pass := "123456"
	isLogin := Login(user, pass)
	// 登录失败
	if !isLogin {
		t.Errorf("user=%s,pass=%s, 用户名必须是root,密码必须是123456", user, pass)
	}
}

// 基准测试
func BenchmarkLogin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		user := "root" + fmt.Sprintf("%d", i)
		pass := "12345" + fmt.Sprintf("%d", i)
		Login(user, pass)
	}
}
