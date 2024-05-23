package queue

import "fmt"

type node struct{
	value int 
	next *node
}

type Queue struct{
	first *node
	last *node
	size int
}
func New() *Queue{
	l := Queue{}
	return &l
}

func (q *Queue) Size() int{
	return q.size
}

func (q *Queue) IsEmpty() bool{
	return q.size == 0
}


func (q *Queue) Add(value int){
	newNode := node{value,nil}
	if q.IsEmpty() {
		q.first = &newNode
		q.last = &newNode
		q.size++
		return
	}
	q.size++
	q.last.next = &newNode
	q.last = q.last.next
}


func (q *Queue) Remove() (int, error) {
	if q.IsEmpty() {
		return  0, fmt.Errorf("queue is empty")
	}
	old := q.first
	q.first = q.first.next
	q.size--
	return old.value, nil
}

func (q * Queue) Print(){
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


func (s *Queue) Peek() (int, error) {
	if s.IsEmpty() {
		return  0, fmt.Errorf("queue is empty")
	}
	return s.first.value, nil
}

