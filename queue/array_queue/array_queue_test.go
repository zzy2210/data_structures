package arrayqueue_test

import (
	arrayqueue "data_structures/queue/array_queue"
	"testing"
)

func TestPush(t *testing.T) {
	q := arrayqueue.New()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	if get := q.Size(); get != 3 {
		t.Errorf("Got %v expected %v", get, 3)
	}
	q.Push(4)
	if get := len(q.Array); get != 8 {
		t.Errorf("Got %v expected %v", get, 8)
	}
	if get := q.QueueBack; get != 4 {
		t.Errorf("Got %v expected %v", get, 4)
	}
	if get := q.Size(); get != 4 {
		t.Errorf("Got %v expected %v", get, 4)
	}
}
