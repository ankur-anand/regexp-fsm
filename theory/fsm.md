## Finite Automata

State Machines has a set of states and transitions between these states.

**So it's a graph with nodes and edges**.

Regular Grammar

S -> bA
A ->  ε

State transition
-> (S)         -b->      ((A))
Staring State            Accepting State

**NFA** NonDeterministic FA -allows transition on the same symbol to different states.

NFA = (Q, ∑, ∆, q0, F)
Q - all possible states - {A,B}
∑ - alphabet - {1, 0}
∆ - transition fn
q0- staring state
F - set of accepting states.