package structure

type MyStack struct {
	data      []interface{}
	size, top int
}

func NewStack(size int) *MyStack {
	return &MyStack{data: make([]interface{}, size), size: size, top: 0}
}

func (stack *MyStack) Length() int {
	return stack.top
}

func (stack *MyStack) Size() int {
	return stack.size
}

func (stack *MyStack) Full() bool {
	return stack.top >= stack.size
}

func (stack *MyStack) Empty() bool {
	return stack.top == 0
}

func (stack *MyStack) Push(value interface{}) bool {
	if stack.Full() {
		return false
	}
	stack.data[stack.top] = value
	stack.top++
	return true
}

func (stack *MyStack) Pop() (res interface{}, ok bool) {
	if stack.Empty() {
		return nil, false
	}
	res = stack.data[stack.top-1]
	stack.top--
	return res, true
}

func (stack *MyStack) Top() (res interface{}, ok bool) {
	if stack.Empty() {
		return nil, false
	}
	return stack.data[stack.top-1], true
}
