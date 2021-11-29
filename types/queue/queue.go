package queue

const blockSize = 4096

type Queue struct {
	tailBlockIdx int
	headBlockIdx int
	tailElemIdx  int
	headElemIdx  int
	blocks       [][]interface{}
	head         []interface{}
	tail         []interface{}
}

func New() *Queue {
	result := new(Queue)
	result.blocks = [][]interface{}{make([]interface{}, blockSize)}
	result.head = result.blocks[0]
	result.tail = result.blocks[0]
	return result
}

func (q *Queue) Push(data interface{}) {
	q.tail[q.tailElemIdx] = data
	q.tailElemIdx++
	if q.tailElemIdx == blockSize {
		q.tailElemIdx = 0
		q.tailBlockIdx = (q.tailBlockIdx + 1) % len(q.blocks)

		if q.tailBlockIdx == q.headBlockIdx {
			buffer := make([][]interface{}, len(q.blocks)+1)
			copy(buffer[:q.tailBlockIdx], q.blocks[:q.tailBlockIdx])
			buffer[q.tailBlockIdx] = make([]interface{}, blockSize)
			copy(buffer[q.tailBlockIdx+1:], q.blocks[q.tailBlockIdx:])
			q.blocks = buffer
			q.headBlockIdx++
			q.head = q.blocks[q.headBlockIdx]
		}
		q.tail = q.blocks[q.tailBlockIdx]
	}
}

func (q *Queue) Pop() interface{} {
	if q.Empty() {
		return nil
	}
	var res interface{}
	res, q.head[q.headElemIdx] = q.head[q.headElemIdx], nil
	q.headElemIdx++
	if q.headElemIdx == blockSize {
		q.headElemIdx = 0
		q.headBlockIdx = (q.headBlockIdx + 1) % len(q.blocks)
		q.head = q.blocks[q.headBlockIdx]
	}
	return res
}

func (q *Queue) Head() interface{} {
	if q.Empty() {
		return nil
	}
	return q.head[q.headElemIdx]
}

func (q *Queue) Tail() interface{} {
	if q.Empty() {
		return nil
	}
	return q.tail[q.tailElemIdx]
}

func (q *Queue) Empty() bool {
	return q.headBlockIdx == q.tailBlockIdx && q.headElemIdx == q.tailElemIdx
}

func (q *Queue) Size() int {
	if q.tailBlockIdx > q.headBlockIdx {
		return (q.tailBlockIdx-q.headBlockIdx)*blockSize - q.headElemIdx + q.tailElemIdx
	} else if q.tailBlockIdx < q.headBlockIdx {
		return (len(q.blocks)-q.headBlockIdx+q.tailBlockIdx)*blockSize - q.headElemIdx + q.tailElemIdx
	} else {
		return q.tailElemIdx - q.headElemIdx
	}
}

func (q *Queue) Reset() {
	*q = *New()
}
