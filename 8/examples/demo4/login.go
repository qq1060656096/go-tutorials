// examples/demo4/login.go
package demo4

// Login 登录
func Login(user, pass string) bool {
	if user == "root" && pass == "123456" {
		return true
	}
	return false
}