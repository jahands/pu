package main

var (
	problemTemplate = `/*
{{problem_description}}
*/
package main

import "github.com/jacobhands/pu"

var (
	answer     = "NA" // Change to correct answer once solved.
	attempts   = 1	  // Number of times answer was submitted to projecteuler.net to get it correct.
	dateSolved = "{{date_solved}}"   // Date the problem was solved. (YYYY-MM-DD)
	problem    = pu.Problem{ID: {{problem_id}}, Solver: solveProblem{{problem_id}}, CorrectAnswer: answer, Attempts: attempts, DateSolved: dateSolved}
)

func main() {
	problem.Answer()
}

// This should return the answer in the same formatting as 'answer' is set to above.
func solveProblem{{problem_id}}() string {
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
	problemDateSolvedString  = "{{date_solved}}"
)
