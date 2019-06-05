package demo5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
func TestLogin(t *testing.T) {
	// 保留恢复SendLoginEmail, 避免后续测试混乱
	var oldSendLoginEmail = SendLoginEmail
	defer func() {
		SendLoginEmail = oldSendLoginEmail
	}()

	SendLoginEmail = func(user string) {
		assert.Equal(t, "root", user)
	}

	b := Login("root", "123456")
	assert.Equal(t, true, b)
}