package fastate

// State Finite Automata State
type State struct {
	accepting bool
	// set of edges or transitions to other states
	// each transition has a label
	transitionsMap map[string][]*State
}

// NewState returns a New State
// which can be either accepting or non accepting
func NewState(accepting bool) *State {
	tm := make(map[string][]*State)
	return &State{
		accepting:      accepting,
		transitionsMap: tm,
	}
}

func (s *State) AddTransitionForSymbol(symbol string, state *State) {
	s.transitionsMap[symbol] = append(s.transitionsMap[symbol], state)
}

func (s *State) GetTransitionsForSymbol(symbol string) []*State {
	return s.transitionsMap[symbol]
}
