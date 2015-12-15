package main

import "testing"

func BenchmarkSolveProblem0(b *testing.B) {
	problem.Bench(b)
}

func TestSolveProblem0(t *testing.T) {
	problem.Test(t)
}
