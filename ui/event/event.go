package event

// Event TODO: StopPropagation (event core?)
type Event interface {
	EventImpl()
}
