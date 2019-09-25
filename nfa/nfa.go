package nfa

import "github.com/ankur-anand/regexp-fsm/fastate"

type NState struct {
	inState  *fastate.State
	outState *fastate.State
}

func NewNState(inState, outState *fastate.State) *NState {
	return &NState{
		inState:  inState,
		outState: outState,
	}
}

// Test whether this NFA matches the string
func (nS *NState) Test(str string) bool {
	return nS.inState.Test(str)
}

// Char is a single Char machine
// -> (in) -a-> ((out))
func Char(symbol string) *NState {
	inState := fastate.NewState(false)
	outState := fastate.NewState(true)
	inState.AddTransitionForSymbol(symbol, outState)

	return NewNState(inState, outState)
}
