/*

 */
package main

import "github.com/jacobhands/pu"

var (
	answer  = "NA" // Change to correct answer once solved.
	problem = pu.Problem{
		ID:            1,
		Solver:        solve,
		CorrectAnswer: answer}
)

func main() {
	problem.Answer()
}

func solve() string {
	return ""
}
