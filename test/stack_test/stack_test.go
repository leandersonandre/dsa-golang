package stack_test

import (
	"leandersonandre/dsa/pkg/stack"
	"testing"
)

func TestCreate(t *testing.T){
l := stack.New()
	if &l == nil {
		t.Error("Stack is not created")
	}
}


func TestPush(t *testing.T){
	l := stack.New()
	l.Push(1)
	if l.IsEmpty() {
		t.Error("Stack is not empty")
	}
	if l.Size() != 1 {
		t.Error("Stack size should be 1")
	}
}

func TestPeek(t *testing.T){
	l := stack.New()
	_, err := l.Peek()
	if err == nil {
		t.Error("Peek in empty Stack should return an error")
	}
	l.Push(1)
	value, err := l.Peek()
	if err != nil || value != 1 {
		t.Error("Top value should be 1")
	}
	l.Push(2)
	value, _ = l.Peek()
	if value != 2 {
		t.Error("Top value should be 2")
	}
}


func TestPop(t *testing.T){
	l := stack.New()
	_, err := l.Peek()
	if err == nil {
		t.Error("Pop in empty Stack should return an error")
	}
	l.Push(1)
	l.Push(4)
	l.Push(8)
	l.Push(99)
	value, err := l.Pop()
	if err != nil || value != 99 {
		t.Errorf("Pop value should be 99 but returned %d",value)
	}
	if l.Size() != 3  {
		t.Errorf("Stack size should be 3 but returned %d",l.Size())
	}
	value, err = l.Pop()
	if value != 8 {
		t.Errorf("Pop value should be 8 but returned %d",value)
	}
}