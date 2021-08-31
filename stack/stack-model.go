package stack

import "sync"

//ItemType - The type of item in stack.
type ItemType interface{}

//Stack - Stack of items.
type Stack struct {
	// Slice of type ItemType, it holds items in stack.
	items []ItemType
	// rwLock for handling concurrent operations on the stack.
	rwLock sync.RWMutex
}
