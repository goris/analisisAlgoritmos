# analisisAlgoritmos

Exercises and assignments for an algorithms analysis class, mostly written in Go.

## Structure

- One Go module for the whole repo (`go.mod` at the root).
- One folder per exercise, each an independent `package main` with its own
  `_test.go` (and optionally `_bench_test.go`).
- No nested `go.mod` files — every exercise shares the root module, so it can
  be built, run, and tested either on its own or together with the rest.

```
analisisAlgoritmos/
├── go.mod
├── Makefile
├── expoBySqr/
│   ├── expoBySqr.go
│   └── expoBySqr_test.go
└── fibonacciMatrix/
    ├── fibonnaci.go
    ├── fibonacci_test.go
    └── fibonacci_bench_test.go
```

## Adding a new exercise

Create a new folder at the root with a `package main` and a `func main()` —
no `go.mod` needed. It's picked up automatically by every command below.

## Usage

Via `make` (recommended):

```bash
make list             # list available exercises
make run EX=expoBySqr # run a single exercise
make test             # run every exercise's tests
make bench            # run every exercise's benchmarks
make build            # build every exercise into bin/<exercise>
make clean            # remove bin/
```

Equivalent raw `go` commands, scoped to a single exercise:

```bash
go run ./expoBySqr
go test ./expoBySqr/...
go test -bench=. -run=^$ ./fibonacciMatrix/...
go build -o bin/expoBySqr ./expoBySqr
```

Or across the whole repo:

```bash
go test ./...
```

Note: plain `go build ./...` from the root fails — Go tries to write the
output binary as e.g. `./fibonacciMatrix`, which collides with the directory
of the same name. Use `make build` (or `go build -o bin/ ./...`) instead.
