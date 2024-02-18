package Stack

import "fmt"

type Stack[T any] struct {
	Top     int
	Element []T
}

func (s *Stack[T]) Push(value T) {
	s.Element = append(s.Element, value)
	s.Top++
}

func (s *Stack[T]) Pop() (*T, error) {
	if s.Top == 0 {
		err := fmt.Errorf("stack is empty")
		return nil, err
	}
	element := s.Element[s.Top-1]
	s.Top--
	return &element, nil
}

func (s *Stack[T]) Peak() {
	if s.Top == 0 {
		fmt.Println("stack is empty")
		return
	}
	fmt.Println(s.Element[s.Top-1])
}

func (s *Stack[T]) Size() int {
	return s.Top
}
