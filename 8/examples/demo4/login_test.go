package demo4

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	b := Login("root", "123456")
	// 断言
	assert.Equal(t, b, true)
}