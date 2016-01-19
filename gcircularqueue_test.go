package gcircularqueue

import (
	"testing"
)

func TestQueueInit(t *testing.T) {
	cq := NewCircularQueue(5)
	if cq.Size != 5 {
		t.Error("It has wrong size:", cq.Size)
	}

	if !cq.IsEmpty() {
		t.Error("It is not empty")
	}
}

func TestPush(t *testing.T) {
	cq := NewCircularQueue(5)
	firstElement := "First"
	cq.Push(firstElement)
	if cq.Elements[0] != firstElement {
		t.Error("It's Push Func is wrong, the first elment is:", cq.Elements[0])
	}
}

func TestShift(t *testing.T) {
	cq := NewCircularQueue(5)
	firstElement := "First"
	cq.Push(firstElement)
	e := cq.Shift()
	if e != firstElement {
		t.Error("It can not shift the first element")
	}
}

func TestIsEmpty(t *testing.T) {
	cq := NewCircularQueue(5)
	if !cq.IsEmpty() {
		t.Error("It's IsEmpty Func is wrong")
	}
}

func TestIsFull(t *testing.T) {
	cq := NewCircularQueue(2)
	cq.Push("First")
	if !cq.IsFull() {
		t.Error("It's IsFull is wrong")
	}
}

func TestFIFO(t *testing.T) {
	cq := NewCircularQueue(3)
	cq.Push(1)
	cq.Push(2)
	firstElement := cq.Shift()
	if firstElement != 1 {
		t.Error("It doesn't support FIFO")
	}
}

func TestCirculartAility(t *testing.T) {
	cq := NewCircularQueue(3)
	cq.Push(1)
	cq.Push(2)
	cq.Shift()
	cq.Push(3)
}