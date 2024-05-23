package linkedlist_test

import (
    "leandersonandre/dsa/pkg/linkedlist"
    "testing"
)

func TestCreate(t *testing.T){
	l := linkedlist.New()
	if &l == nil {
		t.Error("LinkedList is not created")
	}
}

func TestAddFirst(t *testing.T){
	l := linkedlist.New()
	l.AddFirst(1)
	if l.IsEmpty() {
		t.Error("LinkedList is not empty")
	}
}

func TestGetFirst(t *testing.T){
	l := linkedlist.New()
	_, err := l.GetFirst()
	if err == nil {
		t.Error("GetFirst in empty LinkedList should return an error")
	}
	l.AddFirst(1)
	value, err := l.GetFirst()
	if err != nil || value != 1 {
		t.Error("First value should be 1")
	}
	l.AddFirst(2)
	value, _ = l.GetFirst()
	if value != 2 {
		t.Error("First value should be 2")
	}
}


func TestGetLast(t *testing.T){
	l := linkedlist.New()
	_, err := l.GetLast()
	if err == nil {
		t.Error("GetLast in empty LinkedList should return an error")
	}
	l.AddFirst(1)
	value, err := l.GetLast()
	if err != nil || value != 1 {
		t.Error("Last value should be 1")
	}
	l.AddFirst(2)
	value, _ = l.GetLast()
	if value != 1 {
		t.Error("Last value should be 2")
	}
	l.AddLast(3)
	value, _ = l.GetLast()
	if value != 3 {
		t.Error("Last value should be 3")
	}
}

func TestGet(t *testing.T){
	l := linkedlist.New()
	_, err := l.Get(0)
	if err == nil {
		t.Error("Get in empty LinkedList should return an error")
	}
	l.AddFirst(1)
	value, err := l.Get(10)
	if err == nil {
		t.Error("Get in invalid index should return an error")
	}
	l.AddFirst(1)
	l.AddFirst(2)
	l.AddFirst(3)
	l.AddFirst(4)
	value, _ = l.Get(2)
	if value != 2 {
		t.Error("Get(index: 2) value should be 2")
	}
}


func TestRemoveFirst(t *testing.T){
	l := linkedlist.New()
	_, err := l.RemoveFirst()
	if err == nil {
		t.Error("RemoveFirst in empty LinkedList should return an error")
	}
	l.AddFirst(1)
	l.AddFirst(3)
	l.AddFirst(2)
	value, err := l.RemoveFirst()
	if value != 2 {
		t.Error("RemoveFirst should return 2")
	}
	value, err = l.GetFirst()
	if value != 3 {
		t.Error("GetFirst should return 3")
	}
	if l.Size() != 2 {
		t.Error("LinkedList length should be 2")
	}
}


func TestRemoveLast(t *testing.T){
	l := linkedlist.New()
	_, err := l.RemoveLast()
	if err == nil {
		t.Error("RemoveLast in empty LinkedList should return an error")
	}
	l.AddFirst(1)
	l.AddFirst(2)
	l.AddFirst(3)
	value, err := l.RemoveLast()
	if value != 1 {
		t.Errorf("RemoveFirst should return 1 but returned %d",value)
	}
	value, err = l.GetLast()
	if value != 2 {
		t.Errorf("GetLast should return 2 but returned %d",value)
	}
	if l.Size() != 2 {
		t.Error("LinkedList length should be 2")
	}
}



func TestRemove(t *testing.T){
	l := linkedlist.New()
	_, err := l.Remove(10)
	if err == nil {
		t.Error("RemoveLast in empty LinkedList should return an error")
	}
	l.AddFirst(1)
	value, err := l.Remove(10)
	if err == nil {
		t.Error("Remove in invalid index should return an error")
	}
	l.AddFirst(2)
	l.AddFirst(3)
	l.AddFirst(4)
	value, err = l.Remove(1)
	if value != 3 {
		t.Errorf("Remove should return 3 but returned %d",value)
	}
	value, err = l.Get(1)
	if value != 2 {
		t.Errorf("Get(1) should return 2 but returned %d",value)
	}
	if l.Size() != 3 {
		t.Error("LinkedList length should be 3")
	}
}


func TestAdd(t *testing.T){
	l := linkedlist.New()
	err := l.Add(10,99)
	if err == nil {
		t.Error("Add in invalid index should return an error")
	}
    _ = l.Add(0, 1)
	value, err := l.GetFirst()
	if value != 1 {
		t.Errorf("GetFirst should be 1 but returned %d",value)
	}
    _ = l.Add(1, 2)
	_ = l.Add(2,2)
	_ = l.Add(1,99)
	l.PrintList()
	value, err = l.Get(1)
	if value != 99 {
		t.Errorf("Get(1) should be 99 but returned %d",value)
	}
	if l.Size() != 4 {
		t.Error("LinkedList length should be 4")
	}
}