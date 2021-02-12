package internal

import "sort"

type Stack []int

var a, b Stack

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) pop() (int, bool) {
	if s.isEmpty() {
		return 0, false
	}
	element := (*s)[len(*s)-1]
	*s = (*s)[:(len(*s) - 1)]
	return element, true
}

func (s *Stack) push(elem int) {
	*s = append(*s, elem)
}

func (s *Stack) peek() (int, bool) {
	if s.isEmpty() {
		return 0, false
	}
	element := (*s)[len(*s)-1]
	return element, true
}

func (s *Stack) pa(b *Stack) bool {
	elem, f := b.pop()
	if !f {
		return false
	}
	s.push(elem)
	return true
}

func (s *Stack) pb(b *Stack) bool {
	elem, f := s.pop()
	if !f {
		return false
	}
	b.push(elem)
	return true
}

func (s *Stack) sa() bool {
	if len(*s) >= 2 {
		elem1, _ := s.pop()
		elem2, _ := s.pop()
		s.push(elem1)
		s.push(elem2)
		return true
	}
	return false
}

func (s *Stack) ra() bool {
	arr := []int{}
	for !s.isEmpty() {
		elem, _ := s.pop()
		arr = append(arr, elem)
	}
	if arr != nil {
		s.push(arr[0])
		for i := len(arr) - 1; i > 0; i-- {
			s.push(arr[i])
		}
		return true
	}
	return false
}

func (s *Stack) rr() bool {
	arr := []int{}
	for !s.isEmpty() {
		elem, _ := s.pop()
		arr = append(arr, elem)
	}
	if arr != nil {
		for i := len(arr) - 2; i >= 0; i-- {
			s.push(arr[i])
		}
		s.push(arr[len(arr)-1])
		return true
	}
	return false
}

func (s *Stack) isSort() bool {
	for i := len(*s) - 1; i > 0; i-- {
		if (*s)[i] > (*s)[i-1] {
			return false
		}
	}
	return true
}

func (s *Stack) middleNumber() int {
	arr := make([]int, len(*s))
	copy(arr, *s)
	sort.Ints(arr)
	if len(arr)/2 == 0 {
		return arr[len(arr)/2-1]
	}
	return arr[len(arr)/2]
}

func (s *Stack) algorithmSortB() string {
	if len(*s) == 2 && (*s)[1] < (*s)[0] {
		return "sb"
	}
	if (*s)[0] == (*s).maxElem() {
		return "rrb"
	}
	avgB := (*s).middleNumber()
	if (*s)[0] > avgB {
		return "rrb"
	}
	if (*s)[len(*s)-1] > avgB {
		if (*s)[len(*s)-1] < (*s)[len(*s)-2] {
			return "sb"
		}
	} else {
		return "rb"
	}

	return ""
}

func (s *Stack) maxElem() int {
	max := (*s)[len(*s)-1]
	for i := len(*s) - 1; i >= 0; i-- {
		if (*s)[i] > max {
			max = (*s)[i]
		}
	}
	return max
}
