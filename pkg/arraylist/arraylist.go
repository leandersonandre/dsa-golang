package arraylist

type ArrayList struct{
	capacity int
	array []int
	size int
}

func NewArrayList() *ArrayList{
	l := ArrayList{size: 0,capacity: 100}
	var a [100]int
	l.array = &a
	return &l
}