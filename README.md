# analisisAlgoritmos

Exercises and assignments for an algorithms analysis class, mostly written in Go.

## Structure

- One Go module for the whole repo (`go.mod` at the root).
- One folder per exercise, each an independent `package main` with its own
  `_test.go` (and optionally `_bench_test.go`).
- No nested `go.mod` files — every exercise shares the root module, so it can
  be built, run, and tested either on its own or together with the rest.
- One Python environment for the whole repo (`pyproject.toml` + `uv.lock` at the
  root), shared by every exercise's Jupyter notebook.

```
analisisAlgoritmos/
├── go.mod
├── Makefile
├── pyproject.toml             # Python dependencies for the notebooks
├── uv.lock                    # exact versions, managed by uv
├── expoBySqr/
│   ├── expoBySqr.go
│   └── expoBySqr_test.go
├── fibonacciMatrix/
│   ├── fibonnaci.go
│   ├── fibonacci_test.go
│   └── fibonacci_bench_test.go
└── quickSort/
    ├── quickSort.go
    ├── quickSort_test.go
    ├── quickSort_bench_test.go
    ├── generator.go
    ├── generator_test.go
    ├── timing_test.go
    ├── quickSort.ipynb        # graphs the results (Python)
    └── output/                # written by the Go program and the notebook
```

## Notes

- `quickSort` picks its pivot uniformly at random with `math/rand/v2`
  (`rand.IntN`) and partitions with the Lomuto scheme (CLRS
  `RANDOMIZED-PARTITION`). The random pivot makes the expected running time
  O(n log n) on every input, including already sorted or reversed ones.
  Known limitation of Lomuto: an array where every element is equal still
  takes O(n²), because all elements fall on the same side of the pivot.

## Notebooks

The Python environment lives at the repo root and is managed with
[uv](https://docs.astral.sh/uv/). Set it up once (and again whenever
`pyproject.toml` or `uv.lock` changes):

```bash
make notebook-env      # same as: uv sync — creates .venv/ with the exact locked versions
```

Then open a notebook in VS Code and pick the `.venv` kernel, or run
`uv run jupyter lab` from the root. Jupyter runs each notebook from its own
folder, so relative paths like `output/times.csv` point inside that exercise.

To add a Python package: `uv add <package>` (updates `pyproject.toml` and `uv.lock`).

## Graphing the quickSort results

The Go program writes every measured sort time to `output/times.csv`, relative
to the folder it runs from. The notebook reads `quickSort/output/times.csv`, so
run the program from inside `quickSort/`:

```bash
cd quickSort
go run .               # writes output/times.csv and output/summary.txt
```

Then open `quickSort/quickSort.ipynb` as described in [Notebooks](#notebooks).

The notebook saves its graphs to `output/time_vs_n.png` (measured times) and
`output/time_vs_n_fits.png` (measured times with fitted k₁·n·log₂n and k₂·n² curves).

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
make notebook-env     # set up the Python environment for the notebooks
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
