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

func (s *Set) Insert(key interface{}) {
	s.elems[key] = struct{}{}
}

func (s *Set) Contains(key interface{}) bool {
	_, ok := s.elems[key]
	return ok
}

func (s *Set) Delete(key interface{}) {
	delete(s.elems, key)
}

func (s *Set) Empty() bool {
	return s.Len() == 0
}
