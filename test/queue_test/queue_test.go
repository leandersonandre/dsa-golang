package queue_test

import (
    "leandersonandre/dsa/pkg/queue"
	"testing"
)

func TestCreate(t *testing.T){
l := queue.New()
	if &l == nil {
		t.Error("Queue is not created")
	}
}


func TestAdd(t *testing.T){
	l := queue.New()
	l.Add(1)
	if l.IsEmpty() {
		t.Error("Queue is not empty")
	}
	if l.Size() != 1 {
		t.Error("Queue size should be 1")
	}
	l.Add(1)
	if l.Size() != 2 {
		t.Error("Queue size should be 2")
	}
}


func TestPeek(t *testing.T){
	l := queue.New()
	l.Print()
	_, err := l.Peek()
	if err == nil {
		t.Error("Peek in empty Stack should return an error")
	}
	l.Add(1)
	value, err := l.Peek()
	if err != nil || value != 1 {
		t.Error("Peek value should be 1")
	}
	l.Add(2)
	l.Print()
	value, _ = l.Peek()
	if value != 1 {
		t.Error("Peek value should be 1")
	}
}


func TestRemove(t *testing.T){
	l := queue.New()
	_, err := l.Peek()
	if err == nil {
		t.Error("Pop in empty Queue should return an error")
	}
	l.Add(1)
	l.Add(4)
	l.Add(8)
	l.Add(99)
	value, err := l.Remove()
	if err != nil || value != 1 {
		t.Errorf("Remove value should be 1 but returned %d",value)
	}
	if l.Size() != 3  {
		t.Errorf("Queue size should be 3 but returned %d",l.Size())
	}
	value, err = l.Remove()
	if value != 4 {
		t.Errorf("Remove value should be 4 but returned %d",value)
	}
	value, err = l.Remove()
	if value != 8 {
		t.Errorf("Remove value should be 8 but returned %d",value)
	}
	value, err = l.Remove()
	if value != 99 {
		t.Errorf("Remove value should be 99 but returned %d",value)
	}
	if !l.IsEmpty() {
		t.Error("queue should be empty")
	}
	l.Add(1)
	value,err = l.Peek()
	if value != 1 {
		t.Errorf("Peek value should be 1 but returned %d",value)
	}
}