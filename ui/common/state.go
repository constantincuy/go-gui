package common

type State[T any] struct {
	Initial  T
	current  T
	listener func(new T)
}

// OnChange allows to listen to state changes of the current state object
func (state *State[T]) OnChange(listener func(new T)) {
	state.listener = listener

	// Calling the listener to represent the current state
	state.listener(state.current)
}

// Revert resets the current state back to the initial state
func (state *State[T]) Revert() {
	state.SetState(state.Initial)
}

// SetState set the current state
func (state *State[T]) SetState(val T) {
	state.current = val
	if state.listener != nil {
		state.listener(val)
	}
}

func (state *State[T]) Get() T {
	return state.current
}

func NewState[T any](val T) State[T] {
	return State[T]{Initial: val, current: val}
}
