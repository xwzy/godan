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
- [x] Set
- [ ] List
- [ ] Heap / Priority Queue
- [ ] Cache (FIFO / LRU / LFU)
- [x] Trie (Thread Safe)
- [ ] Skip List
- [ ] B Tree
- [ ] B+ Tree
- [ ] AVL Tree
- [ ] RBTree

> ⚠️⚠️  Some following examples are use default constructors for simplicity, you should specify parameters by using NewXXX().  

## Bloom Filter
> 🚀 Check out murmur on [github.com/spaolacci/murmur3](github.com/spaolacci/murmur3) for a better hash function

### Simple Bloom Filter for number
```go
import "github.com/xwzy/godan/bloomfilter"

func main() {
    b := bloomfilter.DefaultNumberBloomFilter()
    
    b.Set((uint64)(100))
    b.Exist((uint64)(100)  // true
    b.Exist((uint64)(1124) // false
}
```

### Simple Bloom Filter for string
```go
import "github.com/xwzy/godan/bloomfilter"

func main() {
    b := bloomfilter.DefaultStringBloomFilter()
    
    b.Set("test0001")
    b.Exist("test001") // true
    b.Exist("test100") // false
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
    buffer.Read(3)   // 55, 56, 57
}
```

## Counter
```go
import "github.com/xwzy/godan/counter"
import "log"

func main () {
    c := counter.DefaultCounter()
    // count in-coming query
    c.CountOne()
    // get the query count
    c.GetTotal()
}
```

## Bitmap
```go
import "github.com/xwzy/godan/bitmap"

func main () {
    b := bitmap.NewBitMap(1000000) 
    // set the number of bit position
    b.Set(1532)
    b.Exist(1532) // true
    b.Exist(532)  // false
    
    b.Delete(1532)
    b.Exist(1532) // false)
}
```

## Stack
```go
import "github.com/xwzy/godan/stack"

func main()  {
    s := stack.DefaultStack()
    s.Push(123)
    s.Top()       // 123
    s.Pop()
}
```

## Queue
```go
import "github.com/xwzy/godan/queue"

func main()  {
    q := queue.DefaultQueue()
    q.Push(111)
    q.Push(222)
    q.Front()    // 111
    q.Back()     // 222
    q.Pop()
    q.Pop()
    q.Empty()    // true
}
```

## Deque
```go
import "github.com/xwzy/godan/deque"

func main()  {
    q := deque.DefaultDeque()
    q.PushBack(222)
    q.PushFront(111)
    q.Front()    // 111
    q.Back()     // 222
    q.PopFront()
    q.PopBack()
    q.Empty()    // true
}
```

## Vector
```go
import "github.com/xwzy/godan/vector"

func main()  {
    vec := vector.DefaultVector()
    vec.PushBack(111)
    vec.Back()    // 111
    vec.PushBack(222)
    vec.At(0)     // 111
    vec.At(1)     // 222
    
    vec.Set(0, 333)
    vec.At(0)     // 333
    
    vec.SwapPosition(0, 1)
    vec.At(0)     // 222
    vec.At(1)     // 333
    
    vec.PopBack()
    vec.Empty()   // false
    vec.Clear()
    vec.Empty()   // true
}
```

## Set
```go
package set

import "github.com/xwzy/godan/set"

func main() {
	s1 := DefaultSet()
	s2 := DefaultSet()

	s1.Add(1)
	s2.Add(2)
	s1.Add(3)
	s2.Add(3)

	s1.InterSection(s2).Exist(3) // true
}
```

## Trie
```go
import "github.com/xwzy/godan/trie"

func main() {
    tree := trie.DefaultTrie()
    tree.Insert("hello")
    tree.Search("hello")   // true
    tree.Search("hellooo") // false
    tree.Search("hel")     // false
    tree.StartsWith("hel") // true

    // support Unicode string
    tree.Insert("你好")
    tree.Search("你好")     // true
    tree.Search("你好啊")   // false
    tree.Search("你")      // false
    tree.StartsWith("你")  // true
}
```

## What does godan mean?
'Godan' is pronounced like "狗蛋" (Dog egg) in Chinese, which happens to be the name of my son.

![](doc/img/godan.png)


