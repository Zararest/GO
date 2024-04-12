package queue

import (
	"sync/atomic"
	"unsafe"
)

// The only way to use atomic.CAS is to use unsafe pointers
type Queue[T any] struct {
	Head unsafe.Pointer
	Tail unsafe.Pointer
}

type Node[T any] struct {
	Value T
	Next  unsafe.Pointer
}

func Create[T any]() Queue[T] {
	// Is this a safe way to use unsafe pointers?
	dummyElem := unsafe.Pointer(&Node[T]{Next: nil})
	return Queue[T]{dummyElem, dummyElem}
}

func (queue Queue[T]) Push(value T) {
	node := Node[T]{Value: value, Next: nil}
	for {
		// trying to get tail and its next element
		tail := atomic.LoadPointer(&queue.Tail)
		tailNode := (*Node[T])(tail)
		next := atomic.LoadPointer(&tailNode.Next)
		// if we capture non-end node, then we shouldn't try to change tail.next
		if next == nil {
			// if tail and tail next are still comsistent,
			//	make tail points to the new node
			// (next should be nil)
			if atomic.CompareAndSwapPointer(&tailNode.Next, next, unsafe.Pointer(&node)) {
				// if this thread managed to insert node, it should update the tail
				// (only if tail hasn't been updated yet)
				atomic.CompareAndSwapPointer(&queue.Tail, tail, unsafe.Pointer(&node))
				return
			}
		} // probably needs else
	}
}
