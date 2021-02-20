package deque

import (
	"github.com/xwzy/godan/assert"
	"testing"
)

func TestQueue(t *testing.T) {
	queue := NewDeque()
	assert.True(t, queue.Empty())

	queue.PushBack(111)

	assert.False(t, queue.Empty())
	assert.Equal(t, queue.Front(), 111)
	assert.Equal(t, queue.Back(), 111)

	queue.PushBack(222)
	assert.Equal(t, queue.Front(), 111)
	assert.Equal(t, queue.Back(), 222)

	queue.PushBack(333)
	assert.Equal(t, queue.Front(), 111)
	assert.Equal(t, queue.Back(), 333)
	assert.Equal(t, queue.Size(), 3)
	assert.False(t, queue.Empty())

	queue.PushFront(444)
	assert.Equal(t, queue.Front(), 444)
	assert.Equal(t, queue.Back(), 333)

	queue.PopBack()
	assert.Equal(t, queue.Front(), 444)
	assert.Equal(t, queue.Back(), 222)

	queue.PopFront()
	assert.Equal(t, queue.Front(), 111)
	assert.Equal(t, queue.Back(), 222)

	queue.PopFront()
	assert.Equal(t, queue.Front(), 222)
	assert.Equal(t, queue.Back(), 222)

	queue.PopFront()
	assert.Equal(t, queue.Front(), nil)
	assert.Equal(t, queue.Back(), nil)

	assert.True(t, queue.Empty())
	assert.Equal(t, queue.Size(), 0)
}
