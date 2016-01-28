// Package pu makes solving projecteuler problems easier and reduces boilerplate.
package pu

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Problem is a Project Euler (.net) problem
type Problem struct {
	// ID is the id of the problem on projecteuler.net
	ID int
	// Solver is the function which solves the problem.
	Solver func() string
	// CorrectAnswer is the answer to the problem (used for testing purposes)
	CorrectAnswer string
	// Attempts is the count of submissions to projecteuler.net it took to submit the correct answer.
	Attempts int
}

// Bench benchmarks the problem. Great for testing for improvements.
func (p Problem) Bench(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p.Solver()
	}
}

// Answer prints out the answer for viewing.
func (p Problem) Answer() {
	if p.CorrectAnswer == "NA" {
		fmt.Println("Problem", p.ID, "has not been solved yet.")
	} else {
		fmt.Println("Answer to problem", p.ID, "is", p.Solver())
	}
}

// Test ensures the answer is correct. Will fail until problem is solved.
func (p Problem) Test(t *testing.T) {
	assert.Equal(t, p.CorrectAnswer, p.Solver(), "These should be equal. Either it hasn't been solved or something broke.")
}
