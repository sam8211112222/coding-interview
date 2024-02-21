package Queue

import "fmt"

type Queue[T any] struct {
	Element []T
}

func (q *Queue[T]) Enqueue(value T) {
	q.Element = append(q.Element, value)
}

// [1:] 表示從索引 1（包括索引 1）一直到切片的末尾。
// [:1] 表示從切片的開始到索引 1（不包括索引 1）。
// [1:3] 表示從索引 1（包括索引 1）到索引 3（不包括索引 3）的元素
func (q *Queue[T]) Dequeue() (*T, error) {
	if len(q.Element) == 0 {
		err := fmt.Errorf("queue is empty")
		return nil, err
	}
	q.Element = q.Element[1:]
	return &q.Element[0], nil
}

func (q *Queue[T]) Peek() (*T, error) {
	if len(q.Element) == 0 {
		return nil, fmt.Errorf("queue is empty")
	}
	return &q.Element[0], nil
}

func (q *Queue[T]) Size() int {
	return len(q.Element)
}
