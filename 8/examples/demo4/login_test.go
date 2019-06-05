package demo4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogin(t *testing.T) {
	b := Login("root", "123456")
	// 断言
	assert.Equal(t, b, true)
}
