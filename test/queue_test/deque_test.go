package queue_test

import (
    "leandersonandre/dsa/pkg/queue"
    "testing"
)

func TestCreateDeque(t *testing.T){
l := queue.NewDeque()
	if &l == nil {
		t.Error("Deque is not created")
	}
}


func TestAddFirst(t *testing.T){
	l := queue.NewDeque()
	l.AddFirst(1)
	if l.IsEmpty() {
		t.Error("Deque is not empty")
	}
	if l.Size() != 1 {
		t.Error("Deque size should be 1")
	}
	l.AddFirst(1)
	if l.Size() != 2 {
		t.Error("Deque size should be 2")
	}
}


func TestGetFirst(t *testing.T){
	l := queue.NewDeque()
	l.Print()
	_, err := l.GetFirst()
	if err == nil {
		t.Error("GetFirst in empty Stack should return an error")
	}
	l.AddFirst(1)
	value, err := l.GetFirst()
	if err != nil || value != 1 {
		t.Error("GetFirst value should be 1")
	}
	l.AddFirst(2)
	l.Print()
	value, _ = l.GetFirst()
	if value != 2 {
		t.Error("GetFirst value should be 2")
	}
}


func TestRemoveFirst(t *testing.T){
	l := queue.NewDeque()
	_, err := l.RemoveFirst()
	if err == nil {
		t.Error("RemoveFirst in empty Deque should return an error")
	}
	l.AddFirst(1)
	l.AddFirst(4)
	l.AddFirst(8)
	l.AddFirst(99)
	value, err := l.RemoveFirst()
	if err != nil || value != 99 {
		t.Errorf("Remove value should be 99 but returned %d",value)
	}
	if l.Size() != 3  {
		t.Errorf("Deque size should be 3 but returned %d",l.Size())
	}
	value, err = l.RemoveFirst()
	if value != 8 {
		t.Errorf("RemoveFirst value should be 8 but returned %d",value)
	}
	l.Print()
	value, err = l.RemoveFirst()
	if value != 4 {
		t.Errorf("RemoveFirst value should be 4 but returned %d",value)
	}
	value, err = l.RemoveFirst()
	if value != 1 {
		t.Errorf("RemoveFirst value should be 1 but returned %d",value)
	}
	if !l.IsEmpty() {
		t.Error("deque should be empty")
	}
	l.AddFirst(1)
	value,err = l.GetFirst()
	if value != 1 {
		t.Errorf("GetFirst value should be 1 but returned %d",value)
	}
}


func TestGetLast(t *testing.T){
	l := queue.NewDeque()
	l.Print()
	_, err := l.GetLast()
	if err == nil {
		t.Error("GetFirst in empty Stack should return an error")
	}
	l.AddFirst(1)
	value, err := l.GetLast()
	if err != nil || value != 1 {
		t.Error("GetLast value should be 1")
	}
	l.AddFirst(2)
	l.Print()
	value, _ = l.GetLast()
	if value != 1 {
		t.Errorf("GetLast value should be 1 but retorned %d",value)
	}
}


func TestRemoveLast(t *testing.T){
	l := queue.NewDeque()
	_, err := l.RemoveLast()
	if err == nil {
		t.Error("RemoveLast in empty Deque should return an error")
	}
	l.AddFirst(1)
	l.AddFirst(4)
	l.AddFirst(8)
	l.AddFirst(99)
	value, err := l.RemoveLast()
	if err != nil || value != 1 {
		t.Errorf("Remove value should be 1 but returned %d",value)
	}
	if l.Size() != 3  {
		t.Errorf("Deque size should be 3 but returned %d",l.Size())
	}
	value, err = l.RemoveLast()
	if value != 4 {
		t.Errorf("RemoveLast value should be 4 but returned %d",value)
	}
	value, err = l.RemoveLast()
	if value != 8 {
		t.Errorf("RemoveLast value should be 8 but returned %d",value)
	}
	value, err = l.RemoveLast()
	if value != 99 {
		t.Errorf("RemoveLast value should be 99 but returned %d",value)
	}
	if !l.IsEmpty() {
		t.Error("deque should be empty")
	}
	l.AddFirst(1)
	value,err = l.GetFirst()
	if value != 1 {
		t.Errorf("GetFirst value should be 1 but returned %d",value)
	}
}
