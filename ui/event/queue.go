package event

type Queue struct {
	listener func(event Event)
}

func (q *Queue) Fire(event Event) {
	if q.listener != nil {
		q.listener(event)
	}
}

func (q *Queue) On(listener func(event Event)) {
	q.listener = listener
}
