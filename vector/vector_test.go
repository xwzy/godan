package vector

import (
	"github.com/xwzy/godan/assert"
	"testing"
)

func TestVector(t *testing.T) {
	vec := DefaultVector()

	vec.PushBack(111)
	assert.Equal(t, vec.Back(), 111)
	assert.Equal(t, vec.Front(), 111)
	assert.False(t, vec.Empty())

	vec.PushBack(222)
	assert.Equal(t, vec.Back(), 222)
	assert.Equal(t, vec.Front(), 111)
	assert.Equal(t, vec.Size(), 2)

	assert.Equal(t, vec.At(0), 111)
	assert.Equal(t, vec.At(1), 222)

	vec.Set(0, 333)
	assert.Equal(t, vec.At(0), 333)

	vec.SwapPosition(0, 1)
	assert.Equal(t, vec.At(0), 222)
	assert.Equal(t, vec.At(1), 333)

	vec.PopBack()
	assert.Equal(t, vec.Back(), 222)
	assert.Equal(t, vec.Front(), 222)
	assert.False(t, vec.Empty())
	vec.PopBack()
	assert.Equal(t, vec.Back(), nil)
	assert.Equal(t, vec.Front(), nil)
	assert.True(t, vec.Empty())

	vec.PushBack(111)
	vec.PushBack(222)
	vec.PushBack(333)
	vec.Clear()
	assert.True(t, vec.Empty())
}
