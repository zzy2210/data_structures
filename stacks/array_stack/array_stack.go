package arraystack

type ArrayStack struct {
}

func (s *ArrayStack) Empty() bool

func (s *ArrayStack) Size() int

func (s *ArrayStack) Top() (val interface{}, ok bool)

func (s *ArrayStack) Pop() (val interface{}, ok bool)

func (s *ArrayStack) Push(val interface{}) (ok bool)
