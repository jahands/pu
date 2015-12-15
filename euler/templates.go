package main

var (
	problemTemplate = `/*
{{problem_description}}*/
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
`
	problemTestTemplate = `package main

import "testing"

func BenchmarkSolveProblem{{problem_id}}(b *testing.B) {
	problem.Bench(b)
}

func TestSolveProblem{{problem_id}}(t *testing.T) {
	problem.Test(t)
}
`
	problemIDString          = "{{problem_id}}"
	problemDescriptionString = "{{problem_description}}"
)
