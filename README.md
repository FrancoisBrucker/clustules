# clustules

A Go library for clustering analysis based on dissimilarity (distance) matrices.

## Overview

`clustules` provides data structures for working with dissimilarity matrices, graphs, and vertex sets, with a focus on hierarchical clustering workflows.

## Packages

- **`vertices`** — labeled vertex set (generic over `string | int`), with bidirectional label ↔ index lookup
- **`diss`** — symmetric dissimilarity matrix (`Diss[T]`); supports reading `.mat` files in lower-triangular, upper-triangular, or square formats, with or without row labels
- **`graph`** — undirected graph (`Graph[T]`) built on a vertex set
- **`modules`** — module/cluster structure grouping multiple graphs (work in progress)

## Usage

```go
import "github.com/FrancoisBrucker/clustules/diss"

data, _ := os.ReadFile("distances.mat")
d, err := diss.NewFromString(string(data))
```

### Matrix file format

`NewFromString` accepts six formats, auto-detected from the token counts per line.

**Lower-triangular with labels** (like `henley.mat`):
```
Ours   0
Chat   47.2  0
Vache  27.7  30.9  0
```

**Square without labels:**
```
0    47.2  27.7
47.2 0     30.9
27.7 30.9  0
```

## Requirements

Go 1.25+

## License

See [LICENSE](LICENSE).
