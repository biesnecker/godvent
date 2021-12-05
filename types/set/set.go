package set

type Set struct {
	elems map[interface{}]struct{}
}

func New() *Set {
	s := new(Set)
	s.elems = make(map[interface{}]struct{})
	return s
}

func (s *Set) Len() int {
	return len(s.elems)
}

func (s *Set) Insert(keys ...interface{}) {
	for _, key := range keys {
		s.elems[key] = struct{}{}
	}
}

func (s *Set) Contains(key interface{}) bool {
	_, ok := s.elems[key]
	return ok
}

func (s *Set) Delete(keys ...interface{}) {
	for _, key := range keys {
		delete(s.elems, key)
	}
}

func (s *Set) Empty() bool {
	return s.Len() == 0
}

func (s *Set) ForEach(cb func(elem interface{})) {
	for k := range s.elems {
		cb(k)
	}
}

func (s *Set) Intersection(other *Set) *Set {
	new := New()
	for k := range s.elems {
		if other.Contains(k) {
			new.Insert(k)
		}
	}
	return new
}
