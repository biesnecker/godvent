package bitvector

import "log"

type BitVector struct {
	capacity int
	elems    []int64
}

func New(cap int) *BitVector {
	needed := cap / 64
	if cap%64 != 0 {
		needed++
	}
	return &BitVector{capacity: cap, elems: make([]int64, needed)}
}

func (bv *BitVector) Set(idx int) {
	if idx < 0 || idx >= bv.capacity {
		log.Fatalln("Index out of range ", idx)
	}
	elem := idx / 64
	mask := int64(1) << (idx % 64)
	bv.elems[elem] |= mask
}

func (bv *BitVector) Unset(idx int) {
	if idx < 0 || idx >= bv.capacity {
		log.Fatalln("Index out of range ", idx)
	}
	elem := idx / 64
	mask := ^(int64(1) << (idx % 64))
	bv.elems[elem] &= mask
}

func (bv *BitVector) Check(idx int) bool {
	if idx < 0 || idx >= bv.capacity {
		log.Fatalln("Index out of range ", idx)
	}
	elem := idx / 64
	mask := int64(1) << (idx % 64)
	return bv.elems[elem]&mask > 0
}

func (bv *BitVector) Clear() {
	for i := range bv.elems {
		bv.elems[i] = int64(0)
	}
}
