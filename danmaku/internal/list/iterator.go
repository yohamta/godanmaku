package list

type Iterator struct {
	list    *List
	current *Element
}

func (ite *Iterator) HasNext() bool {
	return ite.current != nil
}

func (ite *Iterator) Next() *Element {
	e := ite.current
	ite.current = e.GetNext()
	return e
}
