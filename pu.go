// Package pu makes solving projecteuler problems easier and reduces boilerplate.
package pu

import (
	"fmt"
	"testing"
	"time"

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
	// DateSolved is the date the problem was solved.
	DateSolved string
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
		fmt.Println("Problem has not been solved yet.")
	}
	t1 := time.Now()
	answer := p.Solver()
	t2 := time.Now()
	fmt.Println("Problem:", p.ID)
	fmt.Println("Answer:", answer)
	fmt.Println("Execution Time:", t2.Sub(t1))
	fmt.Println("Date Solved:", p.DateSolved)
	fmt.Println("Attempts:", p.Attempts)
}

// Test ensures the answer is correct. Will fail until problem is solved.
func (p Problem) Test(t *testing.T) {
	assert.Equal(t, p.CorrectAnswer, p.Solver(), "These should be equal. Either it hasn't been solved or something broke.")
}
