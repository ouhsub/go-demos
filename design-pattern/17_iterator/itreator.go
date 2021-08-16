package iterator

type Iterator interface {
	HasNext() bool
	Next()
	CurrentItem() interface{}
}

type ArrayInt []int

type ArrayIntIterator struct {
	arrayInt ArrayInt
	index    int
}

func (arr ArrayInt) Iterator() Iterator {
	return &ArrayIntIterator{
		arrayInt: arr,
		index:    0,
	}
}

func (iterator *ArrayIntIterator) HasNext() bool {
	return iterator.index < len(iterator.arrayInt)-1
}

func (iterator *ArrayIntIterator) Next() {
	iterator.index++
}

func (iterator *ArrayIntIterator) CurrentItem() interface{} {
	return iterator.arrayInt[iterator.index]
}
