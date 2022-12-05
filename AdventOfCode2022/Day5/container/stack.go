package container

import "fmt"

type Stack[T comparable] struct {
	vals []T
}

func (s *Stack[T]) Push(val T) {
	s.vals = append([]T{val}, s.vals...)
}

// For some reason, this results in problem in pointers.
func (s *Stack[T]) PushMultiple(val []T) {
	s.vals = append(val, s.vals...)
}

func (s *Stack[T]) Peek() T {
	if len(s.vals) > 0 {
		return s.vals[0]
	}

	var zero T
	return zero
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.vals) == 0
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.vals) == 0 {
		var zero T
		return zero, false
	}
	top := s.vals[0]
	if len(s.vals) <= 1 {
		s.vals = []T{}
	} else {
		s.vals = s.vals[1:len(s.vals)]
	}
	return top, true
}

func (s *Stack[T]) PopMultiple(howManyToPop int) ([]T, bool) {
	if len(s.vals) < howManyToPop {
		return []T{}, false
	}

	top := s.vals[0:howManyToPop]
	if len(s.vals) <= howManyToPop {
		s.vals = []T{}
	} else {
		s.vals = s.vals[howManyToPop:len(s.vals)]
	}

	return top, true
}

func (s Stack[T]) Contains(val T) bool {
	for _, v := range s.vals {
		if v == val {
			return true
		}
	}
	return false
}

func (s *Stack[T]) Print() {
	for i := 0; i < len(s.vals); i++ {
		fmt.Println(s.vals[i])
	}
}
