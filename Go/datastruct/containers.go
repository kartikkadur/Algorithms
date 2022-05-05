package datastruct

// stack data structure
type Stack []int

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(element int) {
	*s = append(*s, element)
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return -1
	}
	index := len(*s)-1
	elem := (*s)[index]
	*s = (*s)[:index]
	return elem
}

// queue datastructure
type Queue []int

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Push(element int) {
	*q = append(*q, element)
}

func (q *Queue) Pop() int {
	if q.IsEmpty() {
		return -1
	}
	element := (*q)[0]
	*q = (*q)[1:]
	return element
}