package linkedlist


import ("fmt")

type Node struct{
	value int 
	next *Node
}

type LinkedList struct{
	head *Node
	length int
}

func New() *LinkedList{
	l := LinkedList{}
	return &l
}

func (l *LinkedList) Length() int{
	return l.length
}
func (l *LinkedList) IsEmpty() bool{
	return l.length == 0
}

func (l *LinkedList) PrintList(){
	fmt.Print("{")
	curr := l.head
	for curr != nil{
		fmt.Print(curr.value)
		if curr.next != nil {
			fmt.Print(", ")
		}
		curr = curr.next
	}
	fmt.Println("}")
}

func (l *LinkedList) AddFirst(value int){
	newNode := Node{value,l.head}
	l.head = &newNode
	l.length++
}


func (l *LinkedList) AddLast(value int){
	if l.IsEmpty() {
		l.AddFirst(value)
		return
	}
	curr := l.head
	for curr.next != nil{
		curr = curr.next
	}
	newNode := Node{value,nil}
	curr.next = &newNode
	l.length++
}


func (l *LinkedList) Add(index int, value int) error {
	if l.Length() < index {
		return  fmt.Errorf("LinkedList: index out of range: {%d}",l.Length())
	}
	if l.IsEmpty() {
		l.AddFirst(value)
		return nil
	}
	if l.Length() == index {
		l.AddLast(value)
		return nil
	}
	curr := l.head.next
	beforeLast := l.head
	count := 1
	for count < index{
		beforeLast = curr
		curr = curr.next
	}
	newNode := Node{value,curr}
	beforeLast.next = &newNode
	l.length++
	return nil
}

func (l *LinkedList) GetFirst() (int, error, ) {
	if l.IsEmpty() {
		return  0, fmt.Errorf("LinkedList is empty")
	}
	return l.head.value, nil 
}

func (l *LinkedList) GetLast() (int, error, ) {
	if l.IsEmpty() {
		return  0, fmt.Errorf("LinkedList is empty")
	}
	curr := l.head
	for curr.next != nil{
		curr = curr.next
	}
	return curr.value, nil 
}


func (l *LinkedList) Get(index int) (int, error, ) {
	if l.IsEmpty() {
		return  0, fmt.Errorf("LinkedList is empty")
	}
	if l.Length() < index {
		return  0, fmt.Errorf("LinkedList: index out of range: {%d}",l.Length())
	}
    count := 0
	curr := l.head
	for count < index{
		curr = curr.next
		count++
	}
	return curr.value, nil 
}


func (l *LinkedList) RemoveFirst() (int, error) {
	if l.IsEmpty() {
		return  0, fmt.Errorf("LinkedList is empty")
	}
	l.length--
	old := l.head
	l.head = l.head.next
	old.next = nil
	return old.value, nil
}


func (l *LinkedList) RemoveLast() (int, error) {
	if l.IsEmpty() {
		return  0, fmt.Errorf("LinkedList is empty")
	}
	if l.Length() == 1 {
		return l.RemoveFirst()
	}
	l.length--
	curr := l.head.next
	beforeLast := l.head
	for curr.next != nil {
		beforeLast = curr
		curr = curr.next
	}
	old := beforeLast.next
	beforeLast.next = nil
	return old.value, nil
}


func (l *LinkedList) Remove(index int) (int, error) {
	if l.IsEmpty() {
		return  0, fmt.Errorf("LinkedList is empty")
	}
	if l.Length() < index {
		return  0, fmt.Errorf("LinkedList: index out of range: {%d}",l.Length())
	}
	if l.Length() == 1 {
		return l.RemoveFirst()
	}
	l.length--
	curr := l.head.next
	beforeLast := l.head
	count := 1
	for count < index  {
		beforeLast = curr
		curr = curr.next
		count++
	}
	beforeLast.next = curr.next
	old := curr
	old.next = nil
	return old.value, nil
}