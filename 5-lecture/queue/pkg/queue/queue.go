package queue

import (
	"errors"
	"sync/atomic"
	"unsafe"
)

// The only way to use atomic.CAS is to use unsafe pointers
type Queue[T any] struct {
	Head unsafe.Pointer // pop from here
	Tail unsafe.Pointer // push here
}

/*
  head.next -> node.next -> tail.next -> nil
	head is always a dummy element
*/

type Node[T any] struct {
	Value T
	Next  unsafe.Pointer
}

func Create[T any]() Queue[T] {
	// Is this a safe way to use unsafe pointers?
	dummyElem := unsafe.Pointer(&Node[T]{Next: nil})
	return Queue[T]{dummyElem, dummyElem}
}

func (queue *Queue[T]) Push(value T) {
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
		} else {
			// if we detect that we have a tail node that is not in queue, we add it
			// (this needs because pop instruction can make previous CAS fail)
			atomic.CompareAndSwapPointer(&queue.Tail, tail, next)
		}
	}
}

func defaultValue[T any]() T {
	var val T
	return val
}

// FIXME: How to return generic default value???!!???
func (queue *Queue[T]) Pop() (T, error) {
	for {
		// trying to get head, tail and its next element
		// WOW: if I change places of head and tail, algorithm won't work
		// 	since compiler isn't allowed to reorder atomic instructions,
		// 	we should firstly capture head, because we firstly check if head == tail
		//	and after changed head
		head := atomic.LoadPointer(&queue.Head)
		tail := atomic.LoadPointer(&queue.Tail)
		next := atomic.LoadPointer(&(*Node[T])(head).Next)
		if head == tail {
			if next == nil {
				return defaultValue[T](), errors.New("queue is empty")
			}
			// if head == tail, probably other thread hasn't made second CAS
			atomic.CompareAndSwapPointer(&queue.Tail, tail, next)
		} else {
			// because head is a dummy
			value := (*Node[T])(next).Value
			// if we still own a head, we return adds new head
			if atomic.CompareAndSwapPointer(&queue.Head, head, next) {
				return value, nil
			}
		}
	}
}

func (queue Queue[T]) Len() int {
	nextPtr := queue.Head
	next := (*Node[T])(nextPtr)
	len := 0
	for next != nil {
		len++
		next = (*Node[T])(next.Next)
	}
	return len
}
