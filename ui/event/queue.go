package event

type Queue struct {
	listeners []func(event Event)
}

func (q *Queue) Fire(event Event) {
	for _, listener := range q.listeners {
		listener(event)
	}
}

func (q *Queue) On(listener func(event Event)) {
	q.listeners = append(q.listeners, listener)
}

func NewEventQueue() Queue {
	return Queue{
		listeners: make([]func(event Event), 0),
	}
}
