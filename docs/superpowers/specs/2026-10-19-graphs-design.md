# Actividad 06 — Grafos: design

## Intent

Class assignment for *Análisis y Diseño de Algoritmos* (Dr. Carlos Villaseñor).
The assignment (`Actividad 06.docx`) asks for two things:

1. **Depth-first search.** Code DFS and use it on `Grafo_no_conexo.txt`, a simple,
   undirected graph that is not connected, to find its connected subgraphs
   (components). Report the vertices of each component, and the code.
2. **Dijkstra.** Code Dijkstra's algorithm and use that same code on
   `Grafo_ponderado.txt`, a weighted **digraph**, to find the minimum route from
   vertex 0 to vertex 14. Report the result, and the code.

The code is graded and shown in a LaTeX report, so it must read like the textbook
algorithm: clear, idiomatic Go with no clever shortcuts. The report is written
later, only after the user signs off on the code. It is out of scope here.

## Goals

- A `graphs/` exercise in the existing repo that, when run, prints both answers.
- DFS that is visibly depth-first: the recursive CLRS `DFS` / `DFS-VISIT` form.
- Dijkstra whose route 0→14 comes from the same function that would solve any
  source/target pair. The route is not hand-picked.
- Tests that check exact answers: vertex lists and the exact path, not just counts
  or the cost.

## Non-goals

- The LaTeX report, figures, and any notebook. They come later and are a separate request.
- Benchmarks or timing. The assignment doesn't ask for them.
- A general graph library: no generics, no string labels, no edge removal.
- Enumerating every shortest path. The program reports the one path Dijkstra finds
  (user decision).

## Anti-goals (would be a failure even if the answers are right)

- A traversal that isn't really depth-first. For example, an iterative stack that
  marks vertices when they are pushed gives the right components but a breadth-ish
  visit order. In an algorithms course that is wrong.
- Answers hard-coded or "assisted" in `main` instead of computed.
- Treating the weighted graph as undirected because the picture has no arrows.

## Constraints

- **Repo conventions** (see `README.md` and `quickSort/`):
  - one root `go.mod` (`module analisisAlgoritmos`, Go 1.27)
  - one folder per exercise, each a `package main` with a `func main()`
  - `make` picks the folder up automatically
- **Test style:** table-driven tests (`tests := []struct{ name string; ... }` with
  `t.Run`), with failure messages shaped like
  `"f(%v) = %v; expected: %v"`.
- **No doc comments** on non-test functions (user preference).
- **Comments in `main` quote the assignment in Spanish**, the way `quickSort/main` does.
- **Standard library only.**
- **Running from the repo root:** `go run ./graphs` (what `make run EX=graphs` does)
  runs from the repo root, so the program must not depend on the working directory
  to find its inputs.

## Input files

The originals are in
`~/Documents/mastersCS/analisisAlgoritmos/actividades/Actividad_06___Grafos/`.
They are copied **verbatim** into `graphs/data/`.

`Grafo_no_conexo.txt` (50 vertices, 84 edges, unweighted):

```
Nodes:
(0, 1, 2, ..., 49)

Edges:

[(0, 29),
 (0, 46),
 ...]
```

`Grafo_ponderado.txt` (20 vertices, 59 edges, weighted):

```
Nodes: 0, 1, 2, ..., 19

Edges:
[(0, 5, {'weight': 1}), (0, 7, {'weight': 9}), ...]
```

The two formats differ in three ways:
- where the vertex list sits: on the `Nodes:` line vs. on the next line in parentheses
- edge tuples with or without `{'weight': w}`
- one edge per line vs. all edges on one line

In both files every edge is written with `u < v`.

## Decided

### Layout

```
graphs/
├── data/
│   ├── Grafo_no_conexo.txt
│   └── Grafo_ponderado.txt
├── graph.go          # graph type + parser
├── graph_test.go
├── dfs.go            # DFS, connected components
├── dfs_test.go
├── dijkstra.go       # Dijkstra + priority queue + path reconstruction
├── dijkstra_test.go
└── main.go           # loads the two files, prints both answers
```

The input files are compiled into the binary with `//go:embed data/*.txt`
(`embed.FS`), so the program runs the same from any directory and the tests read
the same bytes.

### Graph representation

- **Vertices** are the integers `0..n-1`, taken from the `Nodes` section, **not**
  inferred from edges. That way a vertex with no edges still exists and counts as
  its own component.
- **Adjacency list:** `[][]edge`, indexed by vertex, where `edge` is `{to, weight int}`.
- **Directed flag:** a `directed bool` chosen by the caller when parsing.
  - Directed: each tuple `(u, v, w)` adds only `u→v`.
  - Undirected: it adds both `u→v` and `v→u`.
- **Neighbour order:** neighbours are stored in the order edges appear in the file.
  This keeps DFS and Dijkstra deterministic.

### Parser

One function handles both formats. It takes the file text (or an `io.Reader`) plus
the directed flag, and returns the graph or an error.

- **Vertices:** every integer between `Nodes:` and `Edges:`.
- **Edges:** every match of a tuple regex after `Edges:`, `(u, v)` with an optional
  `, {'weight': w}`. It is whitespace and newline tolerant and accepts a sign on `w`
  so negative weights can be detected.
- **Missing weight** means weight 1.
- **Errors:**
  - missing `Nodes:` or `Edges:` section
  - vertex labels that aren't exactly `0..n-1`
  - an edge endpoint outside `0..n-1`
  - a negative weight

### DFS (part 1)

CLRS form: an outer loop over vertices `0..n-1`, and a recursive `dfsVisit` for
each vertex not yet visited.

- Each call to `dfsVisit` from the outer loop discovers exactly one component.
  The function returns `[][]int`, one slice per component, holding the vertices in
  the order DFS discovered them.
- It is run on the **undirected** parse of `Grafo_no_conexo.txt`. The assignment
  calls that graph simple and undirected.
- Recursion depth is at most 50, so recursion is safe.
- **Printing:** `main` prints each component with its size and its vertices
  **sorted ascending**, which is easier to check against the picture.
- **Complexity:** O(V + E).

### Dijkstra (part 2)

- **Signature:** `dijkstra(g, source)` returns `dist []int` and `prev []int`.
  - "Unreachable" is `dist = math.MaxInt` and `prev = -1`.
  - The source itself also has `prev = -1`.
- **Priority queue:** `container/heap` min-heap of `(vertex, distance)` items with
  **lazy deletion**. There is no decrease-key: an improved distance pushes a new
  item, and stale items are skipped when popped. A `visited []bool` marks finalized
  vertices.
- **Relaxation uses strict `<`.** Only a strictly shorter distance replaces the
  current one.
- **Path reconstruction:** `shortestPath(prev, source, target)` follows `prev` back
  from `target` and reverses. It returns `nil` when the target is unreachable, and
  `[source]` when the target is the source.
- It is run on the **directed** parse of `Grafo_ponderado.txt`, because the text
  says *dígrafo* (user decision). The picture has no arrowheads; the report will
  note this.
- **Complexity:** O((V + E) log V).

### Expected answers

Computed with a throwaway script; the tests assert them.

- **Part 1:** 2 components.
  - 35 vertices: `{0, 1, 2, 4, 5, 7, 8, 9, 12, 13, 14, 15, 16, 17, 19, 21, 22, 23, 24,
    25, 26, 27, 28, 29, 31, 32, 34, 38, 39, 40, 41, 43, 44, 46, 48}`
  - 15 vertices: `{3, 6, 10, 11, 18, 20, 30, 33, 35, 36, 37, 42, 45, 47, 49}`
- **Part 2:** the minimum cost 0→14 is **13**, and the path printed is **0 → 5 → 14**
  (1 + 12).
  - **The tie:** `0 → 5 → 8 → 14` (1 + 2 + 10) also costs 13. Strict `<` picks the
    first path deterministically. Vertex 5 is finalized first and sets
    `dist[14] = 13`; when 8 is finalized, 3 + 10 = 13 is not strictly smaller, so
    the earlier path stays. `<=` would flip the answer. The program prints only the
    path Dijkstra returns; the report will mention the tie.

### Output of `go run ./graphs`

Spanish, plain text, stdout only. No `output/` files: the report will quote this
console output. Shape:

```
1. Componentes conexas (DFS) — Grafo_no_conexo.txt
   Número de componentes: 2
   Componente 1 (35 vértices): 0 1 2 4 5 7 ...
   Componente 2 (15 vértices): 3 6 10 11 18 ...

2. Ruta mínima (Dijkstra) — Grafo_ponderado.txt (dígrafo)
   De 0 a 14: 0 -> 5 -> 14
   Costo: 13
```

Errors go to stderr with an `error:` prefix and exit status 1, as in `quickSort`.

### Tests

All table-driven.

**`graph_test.go`**
- both file formats in miniature
- missing weight defaults to 1
- the directed flag adds one adjacency entry per edge; the undirected flag adds two
- each error case: no `Nodes`, no `Edges`, labels not `0..n-1`, endpoint out of range,
  negative weight
- the real files: vertex count 50 / 20, edge count 84 / 59

**`dfs_test.go`**
- a single vertex
- isolated vertices: each is its own component
- one path-shaped component
- several components
- a check that the order is really depth-first: a small graph where DFS and BFS
  visit in different orders, asserting the exact DFS order
- the real file, asserting both sorted vertex lists exactly

**`dijkstra_test.go`**
- source equals target: cost 0, path `[source]`
- an unreachable target: `math.MaxInt`, path `nil`
- a directed edge that can't be walked backwards
- a small graph where the greedy first hop is not the shortest
- a small tie case asserting strict `<` keeps the first-found path
- the real file directed, asserting cost 13 and path `[0 5 14]` exactly
- the real file parsed undirected, asserting the cost is still 13. This documents
  that the missing arrows in the picture don't change the answer.

### Repo housekeeping

- Add `graphs/` to the structure tree in `README.md`.
- Add a short note under **Notes** saying:
  - the weighted file is read as directed
  - the 0→14 tie and how strict `<` resolves it
  - inputs are embedded with `go:embed`
- `graphs/` needs no Makefile change, because new folders are picked up automatically.
  This spec is committed under `docs/`, though, and the Makefile's `EXERCISES` `find`
  would treat `docs` as an exercise and break `make build`. One-line fix: add
  `! -name 'docs'` to that `find`. *(Builder's call, made while writing the spec.)*
- Don't touch the untracked root `output/`.

## Left to the builder

- Exact identifier names, file-internal helper split, and the precise regexes.
- Whether the parser takes a `string` or an `io.Reader`.
- Exact wording and spacing of the Spanish output, within the shape above.
