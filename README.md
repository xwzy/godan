# godan
Data structures implementation in golang.

- [x] Bloom Filter (Thread Safe)
- [x] Ring Buffer (A buffer store the most recent data) (Thread Safe)
- [x] Counter (Count query number in a rate limiter) (Thread Safe)
- [x] Bitmap (Thread Safe)
- [x] Stack
- [x] Queue
- [x] Deque
- [x] Vector
- [ ] List
- [ ] Heap / Priority Queue
- [ ] Cache (FIFO / LRU / LFU)
- [x] Trie (Thread Safe)
- [ ] Skip List
- [ ] B Tree
- [ ] B+ Tree
- [ ] AVL Tree
- [ ] RBTree

> ⚠️⚠️ The following examples are all use default constructors for simplicity, and you should specify parameters by using NewXXX().  

## Bloom Filter
> 🚀 Check out murmur on [github.com/spaolacci/murmur3](github.com/spaolacci/murmur3) for a better hash function

### Simple Bloom Filter for number
```go
import "github.com/xwzy/godan/bloomfilter"

func main() {
    bloomFilter := bloomfilter.DefaultNumberBloomFilter()
    
    bloomFilter.Set((uint64)(100))
    bloomFilter.Exist((uint64)(100) // true
    bloomFilter.Exist((uint64)(1124) // false
}
```

### Simple Bloom Filter for string
```go
import "github.com/xwzy/godan/bloomfilter"

func main() {
    bloomFilter := bloomfilter.DefaultStringBloomFilter()
    
    bloomFilter.Set("test0001")
    bloomFilter.Exist("test001") // true
    bloomFilter.Exist("test100") // false
}
```

## Ring Buffer
```go
import "github.com/xwzy/godan/ringbuffer"

func main() {
    buffer := ringbuffer.NewRingBuffer(10)
    
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
    counter := counter.DefaultCounter()
    
    // count in-coming query
    c.CountOne()
    
    // get the query count
    counter.GetTotal()
}
```

## Bitmap
```go
import "github.com/xwzy/godan/bitmap"

func main () {
    bitmap := bitmap.NewBitMap(1000000) 
    // set the number of bit position
    bitmap.Set(1532)
    bitmap.Exist(1532) // true
    bitmap.Exist(532) // false
    
    bitmap.Delete(1532)
    bitmap.Exist(1532) // false)
}
```

## Trie
```go
import "github.com/xwzy/godan/trie"

func main() {
    tree := trie.DefaultTrie()
    tree.Insert("hello")
    tree.Search("hello") // true
    tree.Search("hellooo") // false
    tree.Search("hel") // false
    tree.StartsWith("hel") //true

    // support Unicode string
    tree.Insert("你好")
    tree.Search("你好") // true
    tree.Search("你好啊") // false
    tree.Search("你") // false
    tree.StartsWith("你") //true
}
```

## What does godan mean?
'Godan' is pronounced like "狗蛋" (Dog egg) in Chinese, which happens to be the name of my son.

![](doc/img/godan.png)


