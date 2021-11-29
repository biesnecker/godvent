package deque

const blockSize = 4096

type Deque struct {
	leftBlockIdx  int
	leftElemIdx   int
	rightBlockIdx int
	rightElemIdx  int

	blocks [][]interface{}
	left   []interface{}
	right  []interface{}
}

func New() *Deque {
	d := new(Deque)
	d.blocks = [][]interface{}{make([]interface{}, blockSize)}
	d.right = d.blocks[0]
	d.left = d.blocks[0]
	return d
}

func (d *Deque) PushRight(data interface{}) {
	d.right[d.rightElemIdx] = data
	d.rightElemIdx++
	if d.rightElemIdx == blockSize {
		d.rightElemIdx = 0
		d.rightBlockIdx = (d.rightBlockIdx + 1) % len(d.blocks)

		if d.rightBlockIdx == d.leftBlockIdx {
			buffer := make([][]interface{}, len(d.blocks)+1)
			copy(buffer[:d.rightBlockIdx], d.blocks[:d.rightBlockIdx])
			buffer[d.rightBlockIdx] = make([]interface{}, blockSize)
			copy(buffer[d.rightBlockIdx+1:], d.blocks[d.rightBlockIdx:])
			d.blocks = buffer
			d.leftBlockIdx++
			d.left = d.blocks[d.leftBlockIdx]
		}
		d.right = d.blocks[d.rightBlockIdx]
	}
}

func (d *Deque) PopRight() (res interface{}) {
	d.rightElemIdx--
	if d.rightElemIdx < 0 {
		d.rightElemIdx = blockSize - 1
		d.rightBlockIdx = (d.rightBlockIdx - 1 + len(d.blocks)) % len(d.blocks)
		d.right = d.blocks[d.rightBlockIdx]
	}
	res, d.right[d.rightElemIdx] = d.right[d.rightElemIdx], nil
	return
}

func (d *Deque) Right() interface{} {
	if d.rightElemIdx > 0 {
		return d.right[d.rightElemIdx-1]
	} else {
		return d.blocks[(d.rightBlockIdx-1+len(d.blocks))%len(d.blocks)][blockSize-1]
	}
}

func (d *Deque) PushLeft(data interface{}) {
	d.leftElemIdx--
	if d.leftElemIdx < 0 {
		d.leftElemIdx = blockSize - 1
		d.leftBlockIdx = (d.leftBlockIdx - 1 + len(d.blocks)) % len(d.blocks)

		if d.leftBlockIdx == d.rightBlockIdx {
			d.leftBlockIdx++
			buffer := make([][]interface{}, len(d.blocks)+1)
			copy(buffer[:d.leftBlockIdx], d.blocks[:d.leftBlockIdx])
			buffer[d.leftBlockIdx] = make([]interface{}, blockSize)
			copy(buffer[d.leftBlockIdx+1:], d.blocks[d.leftBlockIdx:])
			d.blocks = buffer
		}
		d.left = d.blocks[d.leftBlockIdx]
	}
	d.left[d.leftElemIdx] = data
}

func (d *Deque) PopLeft() (res interface{}) {
	res, d.left[d.leftElemIdx] = d.left[d.leftElemIdx], nil
	d.leftElemIdx++
	if d.leftElemIdx == blockSize {
		d.leftElemIdx = 0
		d.leftBlockIdx = (d.leftBlockIdx + 1) % len(d.blocks)
		d.left = d.blocks[d.leftBlockIdx]
	}
	return
}

func (d *Deque) Left() interface{} {
	return d.left[d.leftElemIdx]
}

func (d *Deque) Empty() bool {
	return d.leftBlockIdx == d.rightBlockIdx && d.leftElemIdx == d.rightElemIdx
}

func (d *Deque) Size() int {
	if d.rightBlockIdx > d.leftBlockIdx {
		return (d.rightBlockIdx-d.leftBlockIdx)*blockSize - d.leftElemIdx + d.rightElemIdx
	} else if d.rightBlockIdx < d.leftBlockIdx {
		return (len(d.blocks)-d.leftBlockIdx+d.rightBlockIdx)*blockSize - d.leftElemIdx + d.rightElemIdx
	} else {
		return d.rightElemIdx - d.leftElemIdx
	}
}

func (d *Deque) Reset() {
	*d = *New()
}
