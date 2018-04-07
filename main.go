package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type state struct {
	symbol rune
	edge1  *state
	edge2  *state
}

// helper state
type nfa struct {
	initial *state
	accept  *state
}

// main function
func main() {
	controller := true
	//https://tour.golang.org/flowcontrol/3
	for controller {
		// ask the user which option they want to use
		fmt.Print("\n    Enter \n 1) Infix Expressions Conversion To NFA\n 2) PostFix Expressions Conversion To NFA \n 3) Exit\n")

		// store the users input in a variable
		var input int
		fmt.Scanln(&input)

		// https://tour.golang.org/flowcontrol/9
		// enter switch statement to decide the next option
		switch input {
		case 1:
			fmt.Println("option 1")
			infixToNFA()
		case 2:
			fmt.Println("option 2")
			postFixToNFA()
		case 3:
			fmt.Println("Exiting...")
			// exit the entire application
			controller = false
		default:
			fmt.Println("Please Enter a valid option")
		}
	}
	fmt.Println()
}

// Infix expression conversion to NFA
func infixToNFA() {
	fmt.Print("Please Enter infix expression: ")
	// test (a.(b|c))*
	// error handling the read input
	infixString, err := ReadFromInput()

	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}

	fmt.Println("infix", infixString)

	// convert the infix string to postfix
	postFix := IntoPost(infixString)
	fmt.Println("postfix notation:", postFix)

	fmt.Print("Enter a new String to test if it matches: ")
	newString, err := ReadFromInput()

	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}

	fmt.Println(newString, " matches: ", postFix, " : ", pomatch(postFix, newString))

}

// Postfix Expression conversion to NFA
func postFixToNFA() {
	fmt.Print("Enter postfix expression: ")

	infixString, err := ReadFromInput()

	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	fmt.Print("Enter a new String to test if it matches: ")
	newString, err := ReadFromInput()

	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}

	fmt.Println(newString, " matches: ", newString, " : ", pomatch(infixString, newString))
}

// ReadFromInput reads in user input
func ReadFromInput() (string, error) {

	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')

	return strings.TrimSpace(s), err
}

//post Fix Regular Expression To Non Deterministic Finite Automata
func poregtonfa(pofix string) *nfa {
	// an array of pointers to nfa's that is empty
	nfastack := []*nfa{}

	// looping through each character
	for _, r := range pofix {
		switch r {
		//http://www.fon.hum.uva.nl/praat/manual/Regular_expressions_1__Special_characters.html
		case '.':
			// . the dot matches any character except the newline symbol.
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
			// | the vertical pipe separates a series of alternatives.
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
			// Zero or more
			// pop 1 fragments off the stack of fragments
			frag := nfastack[len(nfastack)-1]
			// up to but not including
			nfastack = nfastack[:len(nfastack)-1]

			accept := state{}
			initial := state{edge1: frag.initial, edge2: &accept}
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &initial})
		case '+':
			// + the plus sign is the match-one-or-more quantifier.
			// pop 1 fragment off nfastack
			frag := nfastack[len(nfastack)-1]
			// create a new state
			accept := state{}
			// new initial accept state
			initial := state{edge1: frag.initial, edge2: &accept}

			// the fragment edge points to initial
			frag.accept.edge1 = &initial

			//Push new fragment to nfastack
			nfastack = append(nfastack, &nfa{initial: frag.initial, accept: &accept})
		case '?':
			// ? the question mark is the match-zero-or-one quantifier.
			// pop 1 fragment off nfastack
			frag := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			initial := state{edge1: frag.initial, edge2: frag.accept}

			nfastack = append(nfastack, &nfa{initial: &initial, accept: frag.accept})
		default:
			// create new accepot, initial state
			accept := state{}
			initial := state{symbol: r, edge1: &accept}

			// push new fragment to the stack
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		}
	}

	// return a single element as a result
	if len(nfastack) != 1 {
		fmt.Println("more than 1 nfa found: length: ", len(nfastack), "nfa Stack: ", nfastack)
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

// IntoPost return to postfix
func IntoPost(infix string) string {
	// mapping the specials in order of importance ()
	specials := map[rune]int{'*': 10, '.': 9, '|': 8, '?': 7, '+': 6}
	postfix := []rune{}
	stack := []rune{}

	for _, r := range infix {
		switch {
		case r == '(':
			stack = append(stack, r)
		case r == ')':
			for stack[len(stack)-1] != '(' {
				postfix = append(postfix, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
		case specials[r] > 0:
			for len(stack) > 0 && specials[r] <= specials[stack[len(stack)-1]] {
				postfix = append(postfix, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, r)
		default:
			postfix = append(postfix, r)
		}
	}
	for len(stack) > 0 {
		postfix = append(postfix, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return string(postfix)
}
