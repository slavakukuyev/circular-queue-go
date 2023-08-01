package main

import (
	"fmt"
	"sync"
)

// CircularQueue represents the circular queue data structure.
type CircularQueue struct {
	data  []interface{}
	front int
	rear  int
	mu    sync.Mutex
}

// NewCircularQueue creates a new circular queue with the given capacity.
func NewCircularQueue(capacity int) *CircularQueue {
	return &CircularQueue{
		data:  make([]interface{}, capacity),
		front: -1,
		rear:  -1,
	}
}

// IsEmpty returns true if the circular queue is empty.
func (cq *CircularQueue) IsEmpty() bool {
	return cq.front == -1
}

// IsFull returns true if the circular queue is full.
func (cq *CircularQueue) IsFull() bool {
	return (cq.rear+1)%len(cq.data) == cq.front
}

// Enqueue adds an element to the rear of the circular queue.
func (cq *CircularQueue) Enqueue(value interface{}) bool {
	cq.mu.Lock()
	defer cq.mu.Unlock()

	if cq.IsFull() {
		return false
	}

	if cq.IsEmpty() {
		cq.front = 0
	}

	cq.rear = (cq.rear + 1) % len(cq.data)
	cq.data[cq.rear] = value
	return true
}

// Dequeue removes the element from the front of the circular queue.
func (cq *CircularQueue) Dequeue() interface{} {
	cq.mu.Lock()
	defer cq.mu.Unlock()

	if cq.IsEmpty() {
		return nil
	}

	value := cq.data[cq.front]

	if cq.front == cq.rear {
		// Queue becomes empty after dequeueing
		cq.front = -1
		cq.rear = -1
	} else {
		cq.front = (cq.front + 1) % len(cq.data)
	}

	return value
}

// Front returns fron pointer and the element at the front of the circular queue without removing it.
func (cq *CircularQueue) Front() (int, interface{}) {
	cq.mu.Lock()
	defer cq.mu.Unlock()

	if cq.IsEmpty() {
		return cq.front, nil
	}

	return cq.front, cq.data[cq.front]
}

// Rear returns rear pointer and the element at the rear of the circular queue without removing it.
func (cq *CircularQueue) Rear() (int, interface{}) {
	cq.mu.Lock()
	defer cq.mu.Unlock()

	if cq.IsEmpty() {
		return cq.rear, nil
	}

	return cq.rear, cq.data[cq.rear]
}

func main() {
	size := 5
	cq := NewCircularQueue(size)

	fmt.Println("IsEmpty:", cq.IsEmpty()) // true

	// Create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	limit := size + 1
	for i := 1; i <= limit; i++ {
		wg.Add(1) // Increment the WaitGroup counter for each goroutine

		// Simulate concurrent enqueue
		go func(i int) {
			defer wg.Done() // Decrement the WaitGroup counter when the goroutine completes
			if cq.Enqueue(i) {
				fmt.Printf("Enqueued: %v\n", i)
			} else {
				fmt.Printf("Failed to enqueue: %v\n", i)
			}
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("IsFull:", cq.IsFull()) // true

	findex, fval := cq.Front()
	fmt.Printf("Front: index=%d, val=%v\n", findex, fval) // Front: index=0, val=[1-6]
	rindex, rval := cq.Rear()
	fmt.Printf("Rear: index=%d, val=%v\n", rindex, rval) // Rear: index=4, val=[1-6]

	for !cq.IsEmpty() {
		fmt.Println("Dequeued:", cq.Dequeue())
	}

	findex, fval = cq.Front()
	fmt.Printf("Front: index=%d, val=%v\n", findex, fval) // Front: index=-1, val=<nil>
	rindex, rval = cq.Rear()
	fmt.Printf("Rear: index=%d, val=%v\n", rindex, rval) // Rear: index=-1, val=<nil>

	cq.Enqueue("bingo")

	findex, fval = cq.Front()
	fmt.Printf("Front: index=%d, val=%v\n", findex, fval) // Front: index=0, val=bingo
	rindex, rval = cq.Rear()
	fmt.Printf("Rear: index=%d, val=%v\n", rindex, rval) // Rear: index=0, val=bingo

}
