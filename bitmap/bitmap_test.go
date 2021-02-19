package bitmap

import (
	"github.com/xwzy/godan/assert"
	"testing"
)

func TestBitMap_Set(t *testing.T) {
	bitmap := NewBitMap(1230600)
	bitmap.Set(1532)

	assert.Equal(t, bitmap.Exist(1532), true)
}

func TestBitMap(t *testing.T) {
	bitmap := NewBitMap(1230600)
	bitmap.Set(1532)

	assert.Equal(t, bitmap.Exist(1532), true)
	assert.Equal(t, bitmap.Exist(234), false)

	bitmap.Set(1533)
	assert.Equal(t, bitmap.Exist(1533), true)
	assert.Equal(t, bitmap.Exist(1532), true)
	assert.Equal(t, bitmap.Exist(1531), false)

	bitmap.Delete(1532)
	assert.Equal(t, bitmap.Exist(1532), false)
	assert.Equal(t, bitmap.Exist(1533), true)

	bitmap.Set(1533)
	assert.Equal(t, bitmap.Exist(1533), true)

	bitmap.Delete(1533)
	assert.Equal(t, bitmap.Exist(1533), false)

	bitmap.Delete(1533)
	assert.Equal(t, bitmap.Exist(1533), false)
}
