package rangeset

type integerRange struct{
	start int
	end int
}

type IntegerRangeSet struct{
	size int
}

func NewIntegerRangeSet() *IntegerRangeSet{
	set := IntegerRangeSet{0}
	return &set
}

func (set *IntegerRangeSet) Add(start int, end int){
	
}

func (set *IntegerRangeSet) Contains(value int) bool{
	return false
}

func (set *IntegerRangeSet) Size(value int) int{
	return set.size
}

func (set *IntegerRangeSet) IsEmpty() bool{
	return set.size == 0
}