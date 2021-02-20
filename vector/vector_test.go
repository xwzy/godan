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

	vec.PopBack()
	assert.Equal(t, vec.Back(), 111)
	assert.Equal(t, vec.Front(), 111)
	assert.False(t, vec.Empty())
	vec.PopBack()
	assert.Equal(t, vec.Back(), nil)
	assert.Equal(t, vec.Front(), nil)
	assert.True(t, vec.Empty())

}
