package deque

type Deque struct {
	data []interface{}
}

func DefaultDeque() *Deque {
	return NewDeque()
}

func NewDeque() *Deque {
	return &Deque{data: make([]interface{}, 0)}
}

func (q *Deque) Empty() bool {
	return q.Size() == 0
}

func (q *Deque) Size() int {
	return len(q.data)
}

func (q *Deque) Front() interface{} {
	if q.Empty() {
		return nil
	}
	return q.data[0]
}

func (q *Deque) Back() interface{} {
	if q.Empty() {
		return nil
	}
	return q.data[len(q.data)-1]
}

func (q *Deque) PushBack(elem interface{}) {
	q.data = append(q.data, elem)
}

func (q *Deque) PushFront(elem interface{}) {
	q.data = append([]interface{}{elem}, q.data...)
}

func (q *Deque) PopFront() {
	if !q.Empty() {
		q.data = q.data[1:]
	}
}

func (q *Deque) PopBack() {
	if !q.Empty() {
		q.data = q.data[:q.Size()-1]
	}
}
