package fastate

import "testing"

func TestState_TransitionForSymbol(t *testing.T) {
	s1 := NewState(false)
	s2 := NewState(true)
	s1.AddTransitionForSymbol("a", s2)
	ts := s1.GetTransitionsForSymbol("a")
	for _, v := range ts {
		t.Log(v)
	}
}
