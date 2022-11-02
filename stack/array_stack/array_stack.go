package arraystack

type ArrayStack struct {
	Array []interface{}
}

func (s *ArrayStack) Empty() bool {
	return len(s.Array) == 0
}

func (s *ArrayStack) Size() int {
	return len(s.Array)
}

func (s *ArrayStack) Top() (val interface{}, ok bool) {
	if s.Empty() {
		return nil, false
	}
	val = s.Array[s.Size()-1]
	return val, true
}

func (s *ArrayStack) Pop() (val interface{}, ok bool) {
	if s.Empty() {
		return nil, false
	}
	val = s.Array[s.Size()-1]
	s.Array = s.Array[:s.Size()-1]
	return val, true
}

func (s *ArrayStack) Push(val interface{}) (ok bool) {
	s.Array = append(s.Array, val)
	return true
}
