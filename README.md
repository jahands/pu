# pu
Project euler helper lib

This is not designed to help solve [Project Euler](https://projecteuler.net) problems. It's more to help standardize the prosess of solving and testing them.

## Getting started:
`go get -u -v github.com/jacobhands/pu/euler`
`euler -new 1 // Creates a folder 'p1' containing 'p1.go' and 'p1_test.go' with starting templates in current directory`
`go test ./... // Test all problems`
`go test -bench=. ./... // Test and benchmark all problems`