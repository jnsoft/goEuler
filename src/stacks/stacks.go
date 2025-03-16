package stacks

// MinStack represents a stack that supports retrieving the minimum value.
// init: ms := &MinStack{}
type MinStack struct {
	stack []pair
}

type pair struct {
	value      int
	currentMin int
}

func NewMinStack() MinStack {
	return MinStack{
		stack: []pair{},
	}
}

// Push adds a new value to the stack and updates the current minimum.
func (ms *MinStack) Push(val int) {
	currentMin := val
	if len(ms.stack) > 0 && ms.stack[len(ms.stack)-1].currentMin < val {
		currentMin = ms.stack[len(ms.stack)-1].currentMin
	}
	ms.stack = append(ms.stack, pair{value: val, currentMin: currentMin})
}

// Pop removes the top value from the stack.
func (ms *MinStack) Pop() {
	if len(ms.stack) > 0 {
		ms.stack = ms.stack[:len(ms.stack)-1]
	}
}

// Top returns the top value of the stack.
func (ms *MinStack) Top() int {
	if len(ms.stack) == 0 {
		panic("stack is empty")
	}
	return ms.stack[len(ms.stack)-1].value
}

// GetMin returns the minimum value in the stack.
func (ms *MinStack) GetMin() int {
	if len(ms.stack) == 0 {
		panic("stack is empty")
	}
	return ms.stack[len(ms.stack)-1].currentMin
}
