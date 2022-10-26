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
	val = s.Array[s.Size()]
	return val, true
}

func (s *ArrayStack) Pop() (val interface{}, ok bool) {
	if s.Empty() {
		return nil, false
	}
	val = s.Array[s.Size()]
	s.Array = s.Array[:s.Size()-1]
	return val, true
}

func (s *ArrayStack) Push(val interface{}) (ok bool) {
	if s.Empty() {
		s.Array = []interface{}{}
	}
	s.Array[s.Size()] = val
	return true
}
