package ringbuffer

import (
	"sync"
	"sync/atomic"
)

// RingBuffer a buffer only remains constant length of the most recent data
// Thread safe with atomic operations
type RingBuffer struct {
	data []interface{}
	pos  uint64
	size uint64
	// this lock is used reversely, read operation should be exclusive and
	// write operation could be parallel
	rwlock sync.RWMutex
}

// DefaultRingBuffer get default ring buffer with size 1024
func DefaultRingBuffer() *RingBuffer {
	return NewRingBuffer(1024)
}

// NewRingBuffer create a new ring buffer with size
func NewRingBuffer(size uint64) *RingBuffer {
	return &RingBuffer{
		data: make([]interface{}, size),
		pos:  0,
		size: size,
	}
}

// Write add new data to ring buffer
func (rb *RingBuffer) Write(data interface{}) bool {
	rb.rwlock.RLock()
	defer rb.rwlock.RUnlock()

	curPos := atomic.LoadUint64(&rb.pos)
	if atomic.CompareAndSwapUint64(&rb.pos, curPos, (curPos+1)%rb.size) {
		rb.data[curPos] = data
		return true
	} else {
		return false
	}
}

// Read read data of given size in ring buffer
// if we have data: 1 2 3 4 5 6 7 8 9, and pos = 3
//        position: 0 1 2 3 4 5 6 7 8
//                        ↑
// we want read size 6, the return data will be
// [7, 8, 9, 1, 2, 3] which is the most recent 6 items
func (rb *RingBuffer) Read(size uint64) []interface{} {
	data := make([]interface{}, 0)
	rb.rwlock.Lock()
	defer rb.rwlock.Unlock()

	curPos := atomic.LoadUint64(&rb.pos)
	if size <= rb.pos {
		data = append(data, rb.data[curPos-size:curPos]...)
	} else {
		remain := size - rb.pos
		data = append(data, rb.data[rb.size-remain:]...)
		data = append(data, rb.data[:curPos]...)
	}
	return data
}

// ReadAll read all data in ring buffer
func (rb *RingBuffer) ReadAll() []interface{} {
	return rb.Read(rb.size)
}
