package bitmap

// BitMap store bool value for key
// uint64 do not need lock or atomic function on 64bit cpu
type BitMap struct {
	size uint64
	data []uint64
}

// DefaultBitMap create a default bitmap with size 1024
func DefaultBitMap() *BitMap {
	return NewBitMap(1024)
}

// NewBitMap create a new bitmap with given size
func NewBitMap(size uint64) *BitMap {
	return &BitMap{
		data: make([]uint64, size>>6+1),
		size: size,
	}
}

// Set set value in bitmap
func (b *BitMap) Set(value uint64) {
	block := value >> 6
	offset := value & 0x3F
	b.data[block] |= 1 << offset
}

// Delete delete value in bitmap
func (b *BitMap) Delete(value uint64) {
	block := value >> 6
	offset := value & 0x3F
	b.data[block] &= ^(1 << offset)
}

// Exist check whether value exist in bitmap
func (b *BitMap) Exist(value uint64) bool {
	block := value >> 6
	offset := value & 0x3F
	return b.data[block]&(1<<offset) != 0
}
