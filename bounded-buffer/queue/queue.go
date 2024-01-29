package queue

type T interface{}

type Queue interface {
	Push(val T)
	Pop() T
	Len() int
}

type queue struct {
	store []T
}

func NewQueue() Queue {
	var store []T
	return &queue{
		store: store,
	}
}

func (q *queue) Push(val T) {
	q.store = append(q.store, val)
}

func (q *queue) Pop() T {
	x := q.store[0]
	q.store = q.store[1:]
	return x
}

func (q *queue) Len() int {
	return len(q.store)
}
