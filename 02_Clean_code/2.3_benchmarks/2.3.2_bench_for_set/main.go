package main

func main() {

}

// реализуйте быстрое множество

type IntSet struct {
	elems *map[int]bool
}

func MakeIntSet() IntSet {
	m := make(map[int]bool)
	return IntSet{&m}
}

func (s *IntSet) Contains(elem int) bool {
	if v := (*s.elems)[elem]; v {
		return true
	}
	return false
}

func (s *IntSet) Add(elem int) bool {
	if s.Contains(elem) {
		return false
	}
	(*s.elems)[elem] = true
	return true
}
