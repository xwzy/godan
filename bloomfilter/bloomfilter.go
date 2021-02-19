package bloomfilter

type BloomFilter struct {
	size uint64
}

func DefaultBloomFilter() {

}

func NewBloomFilter(size uint64, hashFunc func(interface{}) uint64) {

}
