package bitmap

import "log"

// BitMap
type BitMap struct {
	size uint64
	data []byte
}

// DefaultBitMap create a default bitmap with size 1024
func DefaultBitMap() *BitMap {
	return NewBitMap(1024)
}

// NewBitMap create a new bitmap with given size
func NewBitMap(size uint64) *BitMap {
	return &BitMap{
		data: make([]byte, size>>3+1),
		size: size,
	}
}

// Set set value in bitmap
func (b *BitMap) Set(value uint64) {
	block := value >> 3
	offset := value & 0x07
	b.data[block] |= 1 << offset
}

// Delete delete value in bitmap
func (b *BitMap) Delete(value uint64) {
	block := value >> 3
	offset := value & 0x07
	b.data[block] &= ^(1 << offset)
}

// Exist check whether value exist in bitmap
func (b *BitMap) Exist(value uint64) bool {
	block := value >> 3
	offset := value & 0x07
	log.Println(b.data[block] & (1 << offset))
	return b.data[block]&(1<<offset) != 0
}
