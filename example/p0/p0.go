/*
Put problem description here.
*/
package main

import "github.com/jacobhands/pu"

var (
	answer  = "1024" // Change to correct answer once solved.
	problem = pu.Problem{0, solve, answer}
)

func main() {
	problem.Answer()
}

func solve() string {
	return "1024"
}
