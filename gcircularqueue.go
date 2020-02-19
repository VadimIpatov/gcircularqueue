package gcircularqueue

type CircularQueue struct {
	capacity int
	elements []interface{}
	len      int
	head     int
	tail     int
}

// NewCircularQueue creates a new CircularQueue passing an integer as its size
// and returns its pointer
func NewCircularQueue(size int) *CircularQueue {
	cq := CircularQueue{capacity: size + 1, head: 0, tail: 0, len: 0}
	cq.elements = make([]interface{}, cq.capacity)
	return &cq
}

// IsEmpty returns true if the queue is empty
func (c CircularQueue) IsEmpty() bool {
	return c.head == c.tail
}

// IsFull returns true if the queue is full
func (c CircularQueue) IsFull() bool {
	return c.head == (c.tail+1)%c.capacity
}

// Push pushes an element to the queue
// note: if pushing into a full queue, it will panic
func (c *CircularQueue) Push(e interface{}) {
	if c.IsFull() {
		panic("Queue is full")
	}
	c.elements[c.tail] = e
	c.tail = (c.tail + 1) % c.capacity
	c.len++
}

// Shift shifts an element which was pushed the earliest
// note: it will return nil if this queue is empty
func (c *CircularQueue) Shift() (e interface{}) {
	if c.IsEmpty() {
		return nil
	}
	e = c.elements[c.head]
	c.head = (c.head + 1) % c.capacity
	c.len--
	return
}

// Elements returns an underlying buffer's content
func (c *CircularQueue) Elements() []interface{} {
	return c.elements[:c.len]
}

// Len returns a number of elements stored in the queue
func (c *CircularQueue) Len() int {
	return c.len
}
