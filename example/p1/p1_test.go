package main

import "testing"

func BenchmarkSolveProblem1(b *testing.B) {
	problem.Bench(b)
}

func TestSolveProblem1(t *testing.T) {
	problem.Test(t)
}
