/*
Project Euler: Problem 1

Title: Multiples of 3 and 5

Info:
Published on Friday, 5th October 2001, 06:00 pm; Solved by 521276; Difficulty rating: 5%

Description:
If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
Find the sum of all the multiples of 3 or 5 below 1000.
*/
package main

import "github.com/jacobhands/pu"

var (
	answer  = "10" // Change to correct answer once solved. This example has been changed to 10 to prevent build failures.
	problem = pu.Problem{ID: 1, Solver: solveProblem1, CorrectAnswer: answer}
)

func main() {
	problem.Answer()
}

// This should return the answer in the same formatting as 'answer' is set to above.
func solveProblem1() string {
	return "10"
}
