package ringbuffer

import (
	"github.com/xwzy/godan/assert"
	"testing"
)

func TestRingBuffer(t *testing.T) {
	buffer := NewRingBuffer(10)

	for i := 0; i < 5; i++ {
		buffer.Write(i)
	}
	assert.SliceEqual(t, buffer.ReadAll(), []interface{}{nil, nil, nil, nil, nil, 0, 1, 2, 3, 4})

	for i := 5; i < 17; i++ {
		buffer.Write(i)
	}
	assert.SliceEqual(t, buffer.ReadAll(), []interface{}{7, 8, 9, 10, 11, 12, 13, 14, 15, 16})

	for i := 50; i < 58; i++ {
		buffer.Write(i)
	}
	assert.SliceEqual(t, buffer.ReadAll(), []interface{}{15, 16, 50, 51, 52, 53, 54, 55, 56, 57})

	//t.Log(buffer.ReadAll())
}

func TestRingBufferBigSize(t *testing.T) {
	buffer := NewRingBuffer(1024 * 1024)

	for i := 0; i < 1024*1024*10; i++ {
		buffer.Write(i)
	}
	assert.SliceEqual(t, buffer.Read(20), []interface{}{10485740, 10485741, 10485742, 10485743, 10485744, 10485745, 10485746, 10485747, 10485748, 10485749, 10485750, 10485751, 10485752, 10485753, 10485754, 10485755, 10485756, 10485757, 10485758, 10485759})
}
