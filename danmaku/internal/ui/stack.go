package ui

const maxViewNum = 50

type stack struct {
	views [maxViewNum]View
	index int
	ite   *iterator
}

func newStack() *stack {
	s := &stack{}
	s.ite = &iterator{}

	return s
}

func (s *stack) getIterator() *iterator {
	s.ite.stack = s
	s.ite.current = 0
	return s.ite
}

func (s *stack) push(v View) {
	s.index++
	s.views[s.index] = v
}

func (s *stack) pop() {
	s.views[s.index] = nil
	s.index--
}

func (s *stack) size() int {
	return s.index + 1
}
