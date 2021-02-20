package queue

type Queue struct {
	data []interface{}
}

func DefaultQueue() *Queue {
	return NewQueue()
}

func NewQueue() *Queue {
	return &Queue{data: make([]interface{}, 0)}
}

func (q *Queue) Empty() bool {
	return q.Size() == 0
}

func (q *Queue) Size() int {
	return len(q.data)
}

func (q *Queue) Front() interface{} {
	if q.Empty() {
		return nil
	}
	return q.data[0]
}

func (q *Queue) Back() interface{} {
	if q.Empty() {
		return nil
	}
	return q.data[len(q.data)-1]
}

func (q *Queue) Push(elem interface{}) {
	q.data = append(q.data, elem)
}

func (q *Queue) Pop() {
	if !q.Empty() {
		q.data = q.data[1:]
	}
}
