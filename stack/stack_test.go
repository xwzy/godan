package stack

import (
	"github.com/xwzy/godan/assert"
	"testing"
)

func TestNewStack(t *testing.T) {
	stack := DefaultStack()
	assert.True(t, stack.Empty())
	assert.Equal(t, stack.Size(), 0)

	stack.Push(12345)
	assert.Equal(t, stack.Top(), 12345)
	assert.NotEqual(t, stack.Top(), 11111)

	stack.Push(22222)
	assert.Equal(t, stack.Top(), 22222)
	assert.NotEqual(t, stack.Top(), 11111)

	stack.Pop()
	assert.Equal(t, stack.Top(), 12345)
	assert.NotEqual(t, stack.Top(), 11111)
}
