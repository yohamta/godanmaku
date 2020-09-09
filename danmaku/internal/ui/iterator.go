package ui

type iterator struct {
	stack   *stack
	current int
}

func (ite *iterator) hasNext() bool {
	return ite.current < ite.stack.size()
}

func (ite *iterator) next() View {
	ite.current++
	return ite.stack.views[ite.current]
}
