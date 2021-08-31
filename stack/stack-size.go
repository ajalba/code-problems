package stack

func (stack *Stack) Size() int {
	// Acquire read lock
	stack.rwLock.RLock()
	// defer operation of unlock.
	defer stack.rwLock.RUnlock()
	// Return length of items slice.
	return len(stack.items)
}
