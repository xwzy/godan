package set

type Set struct {
	data map[interface{}]struct{}
}

func DefaultSet() *Set {
	return NewSet()
}

func NewSet() *Set {
	return &Set{data: make(map[interface{}]struct{}, 0)}
}

func (s *Set) Add(elem interface{}) {
	s.data[elem] = struct{}{}
}

func (s *Set) Del(elem interface{}) {
	delete(s.data, elem)
}

func (s *Set) Exist(elem interface{}) bool {
	if _, ok := s.data[elem]; ok {
		return true
	}
	return false
}

func (s *Set) InterSection(s1 *Set) *Set {
	res := NewSet()
	for k, _ := range s.data {
		if s1.Exist(k) {
			res.Add(k)
		}
	}
	return res
}

func (s *Set) Union(s1 *Set) *Set {
	res := NewSet()
	for k, _ := range s.data {
		res.Add(k)
	}
	for k, _ := range s1.data {
		res.Add(k)
	}
	return res
}

func (s *Set) Difference(s1 *Set) *Set {
	res := NewSet()
	for k, _ := range s.data {
		if !s1.Exist(k) {
			res.Add(k)
		}
	}
	return res
}

func (s *Set) ToSLice() []interface{} {
	res := make([]interface{}, 0)
	for k, _ := range s.data {
		res = append(res, k)
	}
	return res
}

func (s *Set) ToSLiceUInt64() []uint64 {
	res := make([]uint64, 0)
	for k, _ := range s.data {
		res = append(res, k.(uint64))
	}
	return res
}
