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
