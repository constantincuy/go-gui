package common

import (
	"testing"
)

func TestState_SetState(t *testing.T) {
	s := NewState(0)

	newState := 1
	s.SetState(newState)
	if s.Get() != newState {
		t.Error(IntError("State was not updated", newState, s.Get()))
	}
}

func TestState_Get(t *testing.T) {
	expected := 0
	s := NewState(expected)
	if s.Get() != expected {
		t.Error(IntError("State did not return expected value", expected, s.Get()))
	}
}

func TestState_Revert(t *testing.T) {
	expected := 0
	s := NewState(expected)
	s.SetState(1)
	s.Revert()
	if s.Get() != expected {
		t.Error(IntError("State did not revert to initial value", expected, s.Get()))
	}
}

func TestState_OnChange(t *testing.T) {
	observed := make([]int, 0)
	s := NewState(0)
	s.OnChange(func(new int) {
		observed = append(observed, new)
	})
	s.SetState(1)
	s.Revert()
	if len(observed) != 3 {
		t.Error(IntError("Not all state changes were processed", 3, len(observed)))
	}
}
