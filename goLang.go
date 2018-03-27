package main

import (
	"fmt"
)

type state struct {
	symbol rune
	edge1  *state
	edge2  *state
}

type nfa struct {
	initial *state
	accept  *state
}

//post Fix Regular Expression To Non Deterministic Finite Automata
func poregtonfa(pofix string) *nfa {
	// an array of pointers to nfa's that is empty
	nfastack := []*nfa{}

	// looping through each character
	for _, r := range pofix {
		switch r {
		case '.':
			// pops 2 fragments off the stack of fragments
			frag2 := nfastack[len(nfastack)-1]
			// up to but not including
			nfastack = nfastack[:len(nfastack)-1]

			frag1 := nfastack[len(nfastack)-1]
			// up to but not including
			nfastack = nfastack[:len(nfastack)-1]

			// joins the accept state of the first one to the initial state of the second fragment
			frag1.accept.edge1 = frag2.initial

			// push a new fragment to the stack
			nfastack = append(nfastack, &nfa{initial: frag1.initial, accept: frag2.accept})
		case '|':
			// pops 2 fragments off the stack of fragments
			frag2 := nfastack[len(nfastack)-1]
			// up to but not including
			nfastack = nfastack[:len(nfastack)-1]

			frag1 := nfastack[len(nfastack)-1]
			// up to but not including
			nfastack = nfastack[:len(nfastack)-1]

			// create 2 new states and new initial state and join those two states to the fragment
			accept := state{}
			initial := state{edge1: frag1.initial, edge2: frag2.initial}
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept

			// push a new fragment to the stack
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &initial})
		case '*':
			// pop 1 fragments off the stack of fragments
			frag := nfastack[len(nfastack)-1]
			// up to but not including
			nfastack = nfastack[:len(nfastack)-1]

			accept := state{}
			initial := state{edge1: frag.initial, edge2: &accept}
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &initial})

		default:
			// create new accepot, initial state
			accept := state{}
			initial := state{symbol: r, edge1: &accept}

			// push new fragment to the stack
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &initial})
		}
	}

	// return a single element as a result
	if len(nfastack) != 1 {
		fmt.Println("Error: ", len(nfastack), nfastack)
	}
	return nfastack[0]
}

func addState(l []*state, s *state, a *state) []*state {
	l = append(l, s)

	if s != a && s.symbol == 0 {
		l = addState(l, s.edge1, a)
		if s.edge2 != nil {
			l = addState(l, s.edge2, a)
		}
	}

	return l
}

// does this post fix regualr expression match the string
func pomatch(po string, s string) bool {
	ismatch := false

	// create a post Fix Regular Expression To Non Deterministic Finite Automata from the string po
	ponfa := poregtonfa(po)

	// keeping track of two different states, the current state and the next state
	current := []*state{}
	next := []*state{}

	// pass an array by turning it into a slice.
	current = addState(current[:], ponfa.initial, ponfa.accept)

	// looping through the characters
	for _, r := range s {
		for _, c := range current {
			// check if their labelled by the ruin.
			if c.symbol == r {
				// looping through the current state, check if its equal to the ruine
				next = addState(next[:], c.edge1, ponfa.accept)
			}
		}
		// swap at the end
		current, next = next, []*state{}
	}

	// if any are the accept state, it is a match
	for _, c := range current {
		if c == ponfa.accept {
			ismatch = true
			break
		}
	}

	return ismatch
}

func main() {
	nfa := poregtonfa("ab.c*|")
	fmt.Println(nfa)

	fmt.Println(pomatch("ab.c*|", "cccc"))

}
