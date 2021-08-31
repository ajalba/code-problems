package stack

// Pop removes an Item from the top of the stack
func (stack *Stack) Pop() *ItemType {
	// Checking if stack is empty before performing pop operation
	if len(stack.items) == 0 {
		return nil
	}
	// Acquire read, write lock as items are going to modify.
	stack.rwLock.Lock()
	// Popping item from items slice.
	item := stack.items[len(stack.items)-1]
	//Adjusting the item's length accordingly
	stack.items = stack.items[0 : len(stack.items)-1]
	// Release read write lock.
	stack.rwLock.Unlock()
	// Return last popped item
	return &item
}
