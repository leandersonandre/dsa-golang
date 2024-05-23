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

3. Queue
3. Map

`
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

### Singly Linked List
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

Stack use LIFO (last-in first-out) ordering.

| **Method** | **Complexity** |
|------------|----------------|
| Push(value)     | O(1)           |
| Peek() int   | O(1)           |
| Pop() int   | O(1)           |
| Size() : int | O(1)           |
| IsEmpty() : bool  | O(1)           |

## Queue


Queue use FIFO (first-in first-out) ordering.

| **Method** | **Complexity** |
|------------|----------------|
| Add(value)     | O(1)           |
| Peek() int   | O(1)           |
| Remove() int   | O(1)           |
| Size() : int | O(1)           |
| IsEmpty() : bool  | O(1)           |