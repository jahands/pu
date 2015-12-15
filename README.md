# pu
Project euler helper. Standardize the process of solving PU problems in Go!

## Why?

Project euler contains over 500 problems, so having a standard and quick way of starting a new one is helpful.

## Getting started:

### Install the tool:

`go get -u -v github.com/jacobhands/pu/euler`

#### Create a folder 'p1' containing 'p1.go' and 'p1_test.go' with starting templates in current directory:

`euler -new 1`

([Examples](/example/p0): [p0.go](/example/p0/p0.go), [p0_test.go](/example/p0/p0_test.go))

### Test all problems:

`go test ./...`

### Test and benchmark all problems

`go test -bench=. ./...`