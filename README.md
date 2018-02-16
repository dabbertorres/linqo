# linqo

[![Build Status](https://travis-ci.org/dabbertorres/linqo.svg?branch=master)](https://travis-ci.org/dabbertorres/linqo)
[![Coverage Status](https://coveralls.io/repos/github/dabbertorres/linqo/badge.svg?branch=master)](https://coveralls.io/github/dabbertorres/linqo?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/dabbertorres/linqo)](https://goreportcard.com/report/github.com/dabbertorres/linqo)

I've always thought .NET's LINQ is a pretty nice tool for analyzing data sets - and the LINQ to SQL tool makes it even more useful.

[go-linq](https://github.com/ahmetb/go-linq) is a great library, however, it doesn't support translating to SQL (and isn't the project's goal).

linqo is a library to fill that gap.

linqo is designed with type safety and syntax correctness in mind (at a sacrifice in performance) - which Go's type system makes fairly straightforward to do!

## Getting Started
### Prerequisites
* A working Golang toolchain!

### Installing
```go get -u github.com/dabbertorres/linqo```

(or your favorite Golang vendor/dependency tool)

### Documentation
WIP, see the [package example](linqo_test.go) for now, as well as `go doc` / `godoc` / [godoc.org](godoc.org/github.com/dabbertorres/linqo)

Basic run down of concepts:

As SQL is a (mostly) declarative language, linqo's public API is largely functions, with transparently-used interfaces to guide and direct the method chaining (to most reflect SQL's syntax).

The code may seem to contain a lot of redundant and repetitive interfaces, but there is a method to this madness! While a private API struct will implement several (!) of these interfaces, each function in each interface returns another interface exposing only a much smaller slice of the API. This ensures your Go code, and (assuming linqo is correct) your SQL is all correct at (Go) compile-time. So far, this appears to be working out pretty well!

## Testing
### Prerequisites
* A working Golang toolchain!
### Running
* `go test`
### Example
The [package-wide example](linqo_test.go) depends on [go-sqlite3](github.com/mattn/go-sqlite3) for a temporary in-memory store.

## Versioning, Contributing, etc
linqo is very much an experiment for now - no promises (although SQL syntax hasn't changed much, so the basics are more-or-less stable).

## License
[MIT](LICENSE.md)

## Acknowledgements
[Ron Savage's SQL BNF](https://github.com/ronsavage/SQL)
