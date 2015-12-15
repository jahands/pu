package pu

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Problem is a Project Euler (.net) problem
type Problem struct {
	ID            int
	Solver        func() string
	CorrectAnswer string
}

// Bench benchmarks the problem. Great for testing for improvements.
func (p Problem) Bench(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p.Solver()
	}
}

// Answer prints out the answer for viewing.
func (p Problem) Answer() {
	fmt.Println("Answer to problem " + strconv.Itoa(p.ID) + ": " + p.Solver())
}

// Test ensures that the answer is correct. Will fail until problem is solved.
func (p Problem) Test(t *testing.T) {
	assert.Equal(t, p.CorrectAnswer, p.Solver(), "These should be equal. Either it hasn't been solved or something broke.")
}
