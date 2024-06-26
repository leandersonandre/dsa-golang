# dsa-golang
Data Structures and Algorithms in Go

## Content

1. ArrayList

2. LinkedList

- [x] Singly LinkedList

- [ ] Doubly LinkedList

- [ ] Circular LinkedList

3. Stack

- [x] NodeStack implementation

- [ ] ArrayStack implementation

4. Queue


- [x] Node Queue implementation

- [x] Deque - Double ended queue

6. Binary Heap


10. Sorting algorithms

- [ ] BubleSort

- [ ] MergeSort

- [ ] QuickSort


```
go run .
go test ./...
```

## Array List

| **Method** | **Complexity** |
|------------|----------------|
| get(i)     | O(1)           |
| set(i,e)   | O(1)           |
| add(i,e)   | O(n)           |
| remove(i)  | O(n)           |
| size()     | O(1)           |
| isEmpty()  | O(1)           |

## Linked List

**Singly Linked List** is a list of nodes where each node contains some data and a reference to the next node in the sequence.

| **Method** | **Complexity** |
|------------|----------------|
| AddFirst(value)     | O(1)           |
| AddLast(value)   | O(n)           |
| Add(index,value)   | O(n)           |
| RemoveFirst() : int | O(1)           |
| RemoveLast() : int    | O(n)           |
| Remove(index) : int | O(n)           |
| GetFirst() : int | O(1)           |
| GetLast() : int | O(n)           |
| Get(index) : int | O(n)           |
| PrintList()  | O(n)           |
| Length() : int | O(1)           |
| IsEmpty() : bool  | O(1)           |

## Stack

**Stack** use LIFO (last-in first-out) ordering.

| **Method** | **Complexity** |
|------------|----------------|
| Push(value)     | O(1)           |
| Peek() int   | O(1)           |
| Pop() int   | O(1)           |
| Size() : int | O(1)           |
| IsEmpty() : bool  | O(1)           |

## Queue


**Queue** use FIFO (first-in first-out) ordering.

| **Method** | **Complexity** |
|------------|----------------|
| Add(value)     | O(1)           |
| Peek() int   | O(1)           |
| Remove() int   | O(1)           |
| Size() : int | O(1)           |
| IsEmpty() : bool  | O(1)           |

**Deque** is a queue which elements can be added to or removed from either the front (head) or back (tail). This structure is a implementation of a head-tail doubly linkedlist.


| **Method** | **Complexity** |
|------------|----------------|
| AddFirst(value)     | O(1)           |
| GetFirst() int   | O(1)           |
| RemoveFirst() int   | O(1)           |
| AddLast(value)     | O(1)           |
| GetLast() int   | O(1)           |
| RemoveLast() int   | O(1)           |
| Size() : int | O(1)           |
| IsEmpty() : bool  | O(1)           |


