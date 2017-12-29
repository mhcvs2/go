package structure

type Queue struct {
	data       []interface{}
	size, head, tail int
}

func NewQueue(size int) *Queue {
	return &Queue{
		data:make([]interface{}, size),
		size:size,
		head:0,
		tail:0}
}

func (queue *Queue) Size() int {
	return queue.size
}

func (queue *Queue) Length() int {
	return queue.tail - queue.head
}

func (queue *Queue) Full() bool {
	return queue.tail >= queue.size
}

func (queue *Queue) Empty() bool {
	return queue.tail == queue.head
}

func (queue *Queue) Push(value interface{}) bool {
	if queue.Full() {
		return false
	}
	queue.data[queue.tail] = value
	queue.tail++
	return true
}

func (queue *Queue) Pop() (interface{}, bool) {
	if queue.Empty() {
		return -1, false
	}
	res := queue.data[queue.head]
	queue.head++
	return res, true
}
