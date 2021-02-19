package bloomfilter

import (
	"github.com/xwzy/godan/bitmap"
	"math"
)

// BloomFilter check item existence
type BloomFilter struct {
	size          uint64
	data          *bitmap.BitMap
	hashFunctions []hashFunc
}

// hashFunc hash function type
type hashFunc func(interface{}, uint64) uint64

// DefaultNumberBloomFilter create default number bloom filter with size of 1024
func DefaultNumberBloomFilter() *BloomFilter {
	functionList := make([]hashFunc, 0)
	functionList = append(functionList, func(i interface{}, u uint64) uint64 {
		return i.(uint64) % u
	})
	functionList = append(functionList, func(i interface{}, u uint64) uint64 {
		return uint64(math.Log(float64(i.(uint64)))*float64(2<<10)) % u
	})
	return NewBloomFilter(1024, functionList)
}

// DefaultNumberBloomFilter create default string bloom filter with size of 1024
func DefaultStringBloomFilter() *BloomFilter {
	functionList := make([]hashFunc, 0)
	functionList = append(functionList, func(i interface{}, u uint64) uint64 {
		var sum uint64 = 0
		var index uint64 = 1
		for _, c := range i.(string) {
			sum += uint64(c) * index
			index++
		}
		return sum % u
	})
	functionList = append(functionList, func(i interface{}, u uint64) uint64 {
		var sum uint64 = 0
		var index uint64 = 1
		for _, c := range i.(string) {
			sum += uint64(c) * (1 << index)
			index++
		}
		return sum % u
	})
	return NewBloomFilter(1024, functionList)
}

// NewBloomFilter create new bloom filter
func NewBloomFilter(size uint64, hashFunctions []hashFunc) *BloomFilter {
	return &BloomFilter{
		size:          size,
		data:          bitmap.NewBitMap(size),
		hashFunctions: hashFunctions,
	}
}

// Set set value in bloom filter
func (bf *BloomFilter) Set(value interface{}) {
	for _, f := range bf.hashFunctions {
		bf.data.Set(f(value, bf.size))
	}
}

// Exist check if a value exist in bloom filter
func (bf *BloomFilter) Exist(value interface{}) bool {
	for _, f := range bf.hashFunctions {
		if !bf.data.Exist(f(value, bf.size)) {
			return false
		}
	}
	return true
}
