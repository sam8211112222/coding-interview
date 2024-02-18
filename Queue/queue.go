package Queue

import "fmt"

type Queue[T any] struct {
	Element []T
}

func (q *Queue[T]) Enqueue(value T) {
	q.Element = append(q.Element, value)
}

func (q *Queue[T]) Dequeue() (*T, error) {
	if len(q.Element) == 0 {
		err := fmt.Errorf("queue is empty")
		return nil, err
	}
	return &q.Element[0], nil
}

func (q *Queue[T]) Peak() {
	if len(q.Element) == 0 {
		fmt.Println("queue is empty")
		return
	}
	fmt.Println(q.Element[0])
}

func (q *Queue[T]) Size() int {
	return len(q.Element)
}
