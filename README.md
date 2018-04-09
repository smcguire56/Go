# Graph Theory Go Lang Project

![N|Solid](https://sdtimes.com/wp-content/uploads/2018/02/golang.sh_-490x490.png)

In this project, we are asked to write a program in the Go programming language that can build a non-deterministic finite automaton (NFA) from a regular expression and can use the NFA to check if the regular expression matches any given
string of text.

```go
type nfa struct {
    ...
}
func regexcompile(r string) nfa {
    ...
    return n
}
func (n nfa) regexmatch(n nfa, r sting) bool {
    ...
return ismatch
}
func main() {
    n := regexcompile("01*0")
    t := n.regexmatch("01110")
    f := n.regexmatch("1000001")
}
```
# Instructions

1. Parse the regular expression from infix to postfix notation.
2. Build a series of small NFA’s for parts of the regular expression.
3. Use the smaller NFA’s to create the overall NFA.
4. Implement the matching algorithm using the NFA.


# Marking Scheme
Percent | task| Description
------------ | ------------- | -------------
25%| Research| Investigation of problem and possible solutions.
25%| Development | Clear architecture and well-written code.
25% |Consistency |Good planning and pragmatic attitude to work.
25% |Documentation |Detailed descriptions and explanations.
# References

https://golang.org/ <br />
https://tour.golang.org/flowcontrol/3 <br />
https://tour.golang.org/flowcontrol/9 <br />
http://www.fon.hum.uva.nl/praat/manual/Regular_expressions_1__Special_characters.html <br />
