package listenerstack

type data interface{}

type ListenerStack []data

func (s *ListenerStack) Pop() (item data) {
	if len(*s) == 0 {
		panic("listener expected value: wasnt there")
	}
	item, *s = (*s)[len(*s)-1], (*s)[0:len(*s)-1]
	return item
}

func (s *ListenerStack) Push(d data) {
	*s = append(*s, d)
}
