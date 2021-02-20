package vector

type Vector struct {
	data []interface{}
}

func DefaultVector() *Vector {
	return NewVector()
}

func NewVector() *Vector {
	return &Vector{data: make([]interface{}, 0)}
}

func (q *Vector) Empty() bool {
	return q.Size() == 0
}

func (q *Vector) Size() int {
	return len(q.data)
}

func (q *Vector) Clear() {
	q.data = make([]interface{}, 0)
}

func (q *Vector) Front() interface{} {
	if q.Empty() {
		return nil
	}
	return q.data[0]
}

func (q *Vector) Back() interface{} {
	if q.Empty() {
		return nil
	}
	return q.data[len(q.data)-1]
}

func (q *Vector) PushBack(elem interface{}) {
	q.data = append(q.data, elem)
}

func (q *Vector) PopBack() {
	if !q.Empty() {
		q.data = q.data[:q.Size()-1]
	}
}

func (q *Vector) At(pos int) interface{} {
	return q.data[pos]
}

func (q *Vector) Set(pos int, value interface{}) {
	q.data[pos] = value
}

func (q *Vector) SwapPosition(pos1 int, pos2 int) {
	q.data[pos1], q.data[pos2] = q.data[pos2], q.data[pos1]
}
