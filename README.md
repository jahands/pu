# pu

[![Join the chat at https://gitter.im/jacobhands/pu](https://badges.gitter.im/jacobhands/pu.svg)](https://gitter.im/jacobhands/pu?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
`pu` is a library with a little boilerplate to help with the repetitive creation of new Project Euler projects.

`euler` is a command-line tool to grab the problem from the website and generate a template for each problem as you go.

### Why?

Project euler contains over 500 problems, so having a standard and quick way of starting a new problems is helpful.

One simple command: `euler -new 1` and you're ready to do problem 1 with this:
![[template](https://i.imgur.com/QK3Mk8d.png)](https://i.imgur.com/QK3Mk8dl.png)

Knowing which problems are solved is as easy as `go test ./...` because the template includes a test to see if the problem is solved.

## Getting started:

### Install the tool:

`go get -u -v github.com/jacobhands/pu/euler`

#### Create a folder 'p1' containing 'p1.go' and 'p1_test.go' with starting templates in current directory:

`euler -new 1`

([Examples](/example/p1): [p1.go](/example/p1/p1.go), [p1_test.go](/example/p1/p1_test.go))

### Test all problems:

`go test ./...`

### Test and benchmark all problems

`go test -bench=. ./...`

## Contributing

Contributions are welcome. Please read [CONTRIBUTING.md](/CONTRIBUTING.md) before submitting PR's.
