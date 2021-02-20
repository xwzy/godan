package queue

import (
	"github.com/xwzy/godan/assert"
	"testing"
)

func TestQueue(t *testing.T) {
	queue := NewQueue()
	assert.True(t, queue.Empty())

	queue.Push(112233)

	assert.False(t, queue.Empty())
	assert.Equal(t, queue.Front(), 112233)
	assert.Equal(t, queue.Back(), 112233)

	queue.Push(223344)
	assert.Equal(t, queue.Front(), 112233)
	assert.Equal(t, queue.Back(), 223344)

	queue.Push(334455)
	assert.Equal(t, queue.Front(), 112233)
	assert.Equal(t, queue.Back(), 334455)
	assert.Equal(t, queue.Size(), 3)
	assert.False(t, queue.Empty())

	queue.Pop()
	assert.Equal(t, queue.Front(), 223344)
	assert.Equal(t, queue.Back(), 334455)

	queue.Pop()
	assert.Equal(t, queue.Front(), 334455)
	assert.Equal(t, queue.Back(), 334455)

	queue.Pop()
	assert.Equal(t, queue.Front(), nil)
	assert.Equal(t, queue.Back(), nil)

	assert.True(t, queue.Empty())
	assert.Equal(t, queue.Size(), 0)
}
