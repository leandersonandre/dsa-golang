package stack

import "fmt"

type node struct{
	value int 
	next *node
}

type Stack struct{
	top *node
	size int
}
func New() *Stack{
	l := Stack{}
	return &l
}

func (s *Stack) Size() int{
	return s.size
}

func (s *Stack) IsEmpty() bool{
	return s.size == 0
}

func (s *Stack) Push(value int){
	newNode := node{value,s.top}
	s.top = &newNode
	s.size++
}


func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return  0, fmt.Errorf("stack is empty")
	}
	return s.top.value, nil
}


func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return  0, fmt.Errorf("stack is empty")
	}
	old := s.top
	s.top = s.top.next
	s.size--
	return old.value, nil
}