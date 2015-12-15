# pu
Project euler helper. Standardize the process of solving PU problems in Go!

This won't solve the problems for you, just make it a bit more straightforward.

## Getting started:

### Install the tool:

`go get -u -v github.com/jacobhands/pu/euler`

### Create a folder 'p1' containing 'p1.go' and 'p1_test.go' with starting templates in current directory:

`euler -new 1`

([Examples](/examples/p0)

### Test all problems:

`go test ./...`

### Test and benchmark all problems

`go test -bench=. ./...`