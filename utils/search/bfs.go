package search

import (
	"github.com/biesnecker/godvent/types/queue"
	"github.com/biesnecker/godvent/types/set"
)

type BFSGenerator struct {
	q        *queue.Queue
	seen     *set.Set
	userData interface{}
}

func NewBFSGenerator(initialState Searchable, userData interface{}) *BFSGenerator {
	b := new(BFSGenerator)
	b.q = queue.New()
	b.seen = set.New()
	b.userData = userData

	b.q.Push(initialState)
	b.seen.Insert(initialState.GetRepr(userData))
	return b
}

func (b *BFSGenerator) Next() Searchable {
	if b.q.Empty() {
		return nil
	}
	elem := b.q.Pop().(Searchable)
	for _, next := range elem.GetNext(b.userData) {
		repr := next.GetRepr(b.userData)
		if !b.seen.Contains(repr) {
			b.seen.Insert(repr)
			b.q.Push(next)
		}
	}
	return elem
}

func (b *BFSGenerator) ForEach(handler func(elem Searchable, userData interface{}) bool) {
	for p := b.Next(); p != nil; p = b.Next() {
		if !handler(p, b.userData) {
			break
		}
	}
}
