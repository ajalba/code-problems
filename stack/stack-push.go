package stack

// Push - Adds an Item to the top of the stack
func (stack *Stack) Push(t ItemType) {
	//Initialize items slice if not initialized
	if stack.items == nil {
		stack.items = []ItemType{}
	}
	// Acquire read, write lock before inserting a new item in the stack.
	stack.rwLock.Lock()
	// Performs append operation.
	stack.items = append(stack.items, t)
	// This will release read, write lock
	stack.rwLock.Unlock()
}
