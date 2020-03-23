// examples/demo5/login.go
package demo5

import (
	"fmt"
	"log"
	"net/smtp"
)

const SenderUser = "1060656096@qq.com"
const SenderPass = "123456"
const Hostname = "smtp.qq.com"

// 发送邮件
var SendLoginEmail = func(user string) {
	msg := fmt.Sprintf("%s, Welcome to login", user)
	auth := smtp.PlainAuth("", SenderUser, SenderPass, Hostname)
	if err := smtp.SendMail(Hostname, auth, SenderUser, []string{user}, []byte(msg)); err != nil {
		log.Printf("smtp.SendMail(%s) failed: %s", user, err)
	}
}

// Login 登录
func Login(user, pass string) bool {
	if user == "root" && pass == "123456" {
		SendLoginEmail(user)
		return true
	}
	return false
}
