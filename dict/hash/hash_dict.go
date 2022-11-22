package hash

type HashDict struct {
	Buckets []*chainNode
	size    int
}

type chainNode struct {
	Key  interface{}
	Val  interface{}
	Next *chainNode
}

func NewHashDict() *HashDict {
	return &HashDict{
		Buckets: make([]*chainNode, 11),
	}
}

func (d *HashDict) Empty() bool {
	return d.size == 0
}

func (d *HashDict) Size() int {
	return d.size
}

// 这里使用hash函数为f(k)=k%11
func (d *HashDict) Get(key interface{}) (interface{}, bool) {
	var bucket int
	switch key.(type) {
	case int:
		bucket = key.(int) % 11
	case string:
		var ascii rune
		for _, v := range key.(string) {
			ascii += v
		}
		bucket = int(ascii) % 11
	default:
		// 个人能力有限，仅提供对string与int类型
		return nil, false
	}
	for node := d.Buckets[bucket]; node != nil; node = node.Next {
		if node.Key == key {
			return node.Val, true
		}
	}
	return nil, false
}

func (d *HashDict) Insert(key, val interface{}) bool {
	var bucket int
	switch key.(type) {
	case int:
		bucket = key.(int) % 11
	case string:
		var ascii rune
		for _, v := range key.(string) {
			ascii += v
		}
		bucket = int(ascii) % 11
	default:
		// 个人能力有限，仅提供对string与int类型
		return false
	}
	if d.Buckets[bucket] == nil {
		d.Buckets[bucket] = &chainNode{
			Key: key,
			Val: val,
		}
		d.size++
		return true
	}
	for node := d.Buckets[bucket]; node != nil; node = node.Next {
		if node.Key == key {
			node.Val = val
		}
		if node.Next == nil {
			node.Next = &chainNode{
				Key: key,
				Val: val,
			}
			d.size++
		}
	}
	return true
}

func (d *HashDict) Erase(key interface{}) bool {
	var bucket int
	switch key.(type) {
	case int:
		bucket = key.(int) % 11
	case string:
		var ascii rune
		for _, v := range key.(string) {
			ascii += v
		}
		bucket = int(ascii) % 11
	default:
		// 个人能力有限，仅提供对string与int类型
		return false
	}
	if d.Buckets[bucket] == nil {
		return true
	} else if d.Buckets[bucket].Key == key {
		d.Buckets[bucket] = d.Buckets[bucket].Next
		d.size--
	}
	for node := d.Buckets[bucket]; node.Next != nil; node = node.Next {
		if node.Next.Key == key {
			node.Next = node.Next.Next
			d.size--
			return true
		}
	}
	return true
}
