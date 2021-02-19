# godan
Data structures implementation in golang

## Bloom Filter
### Simple Bloom Filter for number
```go
import "github.com/xwzy/godan/bloomfilter"
func main() {
    bloomFilter := DefaultNumberBloomFilter()
    
    bloomFilter.Set((uint64)(100))
    bloomFilter.Exist((uint64)(100) // true
    bloomFilter.Exist((uint64)(1124) // false
}
```

### Simple Bloom Filter for string
```go
import "github.com/xwzy/godan/bloomfilter"
func main() {
    bloomFilter := DefaultStringBloomFilter()
    
    bloomFilter.Set("test0001")
    bloomFilter.Exist("test001") // true
    bloomFilter.Exist("test100") // false
}
```

## Ring Buffer
```go
import "github.com/xwzy/godan/ringbuffer"
func main() {
    buffer := NewRingBuffer(10)
    
    for i := 0; i < 5; i++ {
        buffer.Write(i)
    }
    buffer.ReadAll() // nil, nil, nil, nil, nil, 0, 1, 2, 3, 4
    
    for i := 5; i < 17; i++ {
        buffer.Write(i)
    }
    buffer.ReadAll() // 7, 8, 9, 10, 11, 12, 13, 14, 15, 16
    
    for i := 50; i < 58; i++ {
        buffer.Write(i)
    }
    buffer.ReadAll() // 15, 16, 50, 51, 52, 53, 54, 55, 56, 57
    buffer.Read(3) // 55, 56, 57
}
```

## Counter
```go
import "github.com/xwzy/godan/counter"
import "log"

func main () {
    counter := DefaultCounter()
    
    // count in coming query
    c.CountOne()
    
    // get the query count
    counter.GetTotal()
}
```

## Bitmap
```go
import "github.com/xwzy/godan/bitmap"
func main () {
    bitmap := NewBitMap(1000000) 
    // set the number of bit position
    bitmap.Set(1532)
    bitmap.Exist(1532) // true
    bitmap.Exist(532) // false
    
    bitmap.Delete(1532)
    bitmap.Exist(1532) // false)
}
```