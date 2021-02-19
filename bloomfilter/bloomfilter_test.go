package bloomfilter

import (
	"github.com/xwzy/godan/assert"
	"testing"
)

func TestBloomFilter(t *testing.T) {

}

func TestDefaultNumberBloomFilter(t *testing.T) {
	bloomFilter := DefaultNumberBloomFilter()

	bloomFilter.Set((uint64)(100))
	assert.False(t, bloomFilter.Exist((uint64)(1124)))
	assert.False(t, bloomFilter.Exist((uint64)(2148)))
	assert.True(t, bloomFilter.Exist((uint64)(100)))
}

func TestDefaultStringBloomFilter(t *testing.T) {
	bloomFilter := DefaultStringBloomFilter()

	bloomFilter.Set("test0001")
	assert.False(t, bloomFilter.Exist("test001"))
	assert.True(t, bloomFilter.Exist("test0001"))

	bloomFilter.Set("test0001")
	assert.False(t, bloomFilter.Exist("test001"))
	assert.True(t, bloomFilter.Exist("test0001"))

	bloomFilter.Set("test2233")
	assert.False(t, bloomFilter.Exist("test2332"))
	assert.False(t, bloomFilter.Exist("test2332"))
	assert.True(t, bloomFilter.Exist("test2233"))
}
