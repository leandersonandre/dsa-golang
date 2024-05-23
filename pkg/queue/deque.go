package queue

import "fmt"

type dequeNode struct{
	value int 
	next *dequeNode
	previous *dequeNode
}

type Deque struct{
	first *dequeNode
	last *dequeNode
	size int
}


func NewDeque() *Deque{
	l := Deque{}
	return &l
}

func (q *Deque) Size() int{
	return q.size
}

func (q *Deque) IsEmpty() bool{
	return q.size == 0
}


func (q *Deque) AddFirst(value int){
	newNode := dequeNode{value,	q.first, nil}
	if q.IsEmpty() {
		q.first = &newNode
		q.last = &newNode
		q.size++
		return
	}
	q.size++
	q.first.previous = &newNode
	q.first = &newNode
}


func (q *Deque) RemoveFirst() (int, error) {
	if q.IsEmpty() {
		return  0, fmt.Errorf("deque is empty")
	}
	q.size--
	old := q.first
	if q.IsEmpty() {
		q.first = nil
		q.last = nil
		return old.value, nil
	}
	q.first = q.first.next
	q.first.previous = nil
	return old.value, nil
}


func (q *Deque) RemoveLast() (int, error) {
	if q.IsEmpty() {
		return  0, fmt.Errorf("deque is empty")
	}
	old := q.last
	q.size--
	if q.IsEmpty() {
		q.first = nil
		q.last = nil
		return old.value, nil
	}
	q.last = q.last.previous
	q.last.next = nil
	return old.value, nil
}


func (s *Deque) GetFirst() (int, error) {
	if s.IsEmpty() {
		return  0, fmt.Errorf("deque is empty")
	}
	return s.first.value, nil
}

func (s *Deque) GetLast() (int, error) {
	if s.IsEmpty() {
		return  0, fmt.Errorf("deque is empty")
	}
	return s.last.value, nil
}


func (q * Deque) Print(){
	fmt.Print("{")
	curr := q.first
	for curr != nil{
		fmt.Print(curr.value)
		if curr.next != nil {
			fmt.Print(", ")
		}
		curr = curr.next
	}
	fmt.Println("}")
}