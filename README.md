# Circular Queue in Go

Circular Queue is a linear data structure implemented in Go based on the First In First Out (FIFO) principle. It supports concurrent Enqueue and Dequeue operations while maintaining a fixed-size circular buffer. The circular queue efficiently reuses memory space and avoids shifting elements when the queue is full.

## Features

- Concurrent Enqueue and Dequeue operations with thread-safe access using `sync.Mutex`.
- Efficient utilization of memory through circular buffer implementation.
- Provides methods to check if the queue is empty or full.
- Easy-to-use API to Enqueue and Dequeue elements from the front and rear of the queue.
