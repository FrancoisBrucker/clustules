# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
go build ./...          # Build all packages
go test ./...           # Run all tests
go test ./diss/...      # Run tests for a single package
go test -run TestName ./diss/  # Run a specific test
go run main.go          # Run the entry point (reads henley.mat)
```

## Architecture

This is a Go library (`github.com/FrancoisBrucker/clustules`) for clustering analysis using dissimilarity matrices. All data structures use Go generics (type parameter `T` constrained to `string | int` for vertex labels).

**Package dependency order:**

1. **`vertices`** — foundation: labeled vertex set, no duplicates. Provides `Label(index)` ↔ `Index(label)` mapping.
2. **`diss`** — builds on vertices: symmetric distance matrix (`Diss[T]`). The main algorithmic package. Parses `.mat` files in six formats (lower/upper/square triangular, with or without row/column labels) via `NewFromString()` and auto-detected `matrixType()`.
3. **`graph`** — builds on vertices: undirected graph (`Graph[T]`) with adjacency stored as `map[int]map[int]bool`.
4. **`modules`** — builds on graph: `Indistinct[T]` groups multiple graphs for module/cluster analysis. Currently incomplete (`NewFromRelation` has no return statement).

The entry point (`main.go`) demonstrates usage: reads `henley.mat` (12-animal French dataset, lower-triangular with labels) into a `Diss[string]`, prints it.

## Key conventions

- `Set(i, j, v)` on `diss.Matrix` enforces symmetry by setting both `[i][j]` and `[j][i]`.
- `Update(i, j, fn)` applies a function to both symmetric positions.
- Matrix indices are integers; labels are resolved through the embedded `vertices.Vertices[T]`.
