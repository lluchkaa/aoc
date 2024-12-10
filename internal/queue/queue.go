package queue

type node[T any] struct {
	value T
	next  *node[T]
}

type Queue[T any] struct {
	head *node[T]
	tail *node[T]
}

func (q *Queue[T]) Push(v T) {
	n := &node[T]{value: v}

	if q.tail != nil {
		q.tail.next = n
	}

	q.tail = n

	if q.head == nil {
		q.head = n
	}
}

func (q *Queue[T]) Pop() (T, bool) {
	if q.head == nil {
		var v T
		return v, false
	}

	v := q.head.value
	q.head = q.head.next

	if q.head == nil {
		q.tail = nil
	}

	return v, true
}

func (q *Queue[T]) Peek() (T, bool) {
	if q.head == nil {
		var v T
		return v, false
	}

	return q.head.value, true
}

func New[T any]() *Queue[T] {
	return &Queue[T]{}
}
