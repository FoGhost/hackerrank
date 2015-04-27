package main
import "fmt"
type cell struct {
  depth int
  isPossibleCavity bool
}
func main() {
  var n int
  fmt.Scan(&n)


  cells_h := make([]*cell, n * n)
  cells_v := make([]*cell, n * n)
  for i := 0; i < n; i++ {
    var row_str string
    fmt.Scan(&row_str)

    for j, depth := range row_str {
      var c cell
      c.depth = int(depth - '0')
      c.isPossibleCavity = true

      if j == 0 {
        // first cell of row
        c.isPossibleCavity = false
      } else {
        pre_cell := cells_h[i * n + j - 1] // get pre cell
        if pre_cell.depth > c.depth {
          c.isPossibleCavity = false
        } else if pre_cell.depth == c.depth {
          pre_cell.isPossibleCavity = false
          c.isPossibleCavity = false
        } else {
          pre_cell.isPossibleCavity = false
        }

        // last cell of row
        if j == (n - 1) {
          c.isPossibleCavity = false
        }
      }

      cells_h[i * n + j] = &c
      cells_v[j * n + i] = &c
    }
  }

  for j := 0; j < n; j++ {
    for i := 0; i < n; i++ {
      current_cell := cells_v[j * n + i]

      if i == 0 {
        // first cell of column
        current_cell.isPossibleCavity = false
      } else {
        current_cell := cells_v[j * n + i]
        pre_cell := cells_v[j * n + i - 1]

        if pre_cell.depth > current_cell.depth {
          current_cell.isPossibleCavity = false
        } else if pre_cell.depth == current_cell.depth {
          pre_cell.isPossibleCavity = false
          current_cell.isPossibleCavity = false
        } else {
          pre_cell.isPossibleCavity = false
        }

        // last cell of column
        if i == (n - 1) {
          current_cell.isPossibleCavity = false
        }
      }
    }
  }

  // output
  for i := 0; i < n; i++ {
    for j := 0; j < n; j++ {
      current_cell := cells_h[i * n + j]
      if current_cell.isPossibleCavity {
        fmt.Printf("%s", "X")
      } else {
        fmt.Printf("%d", current_cell.depth)
      }
      if j == n - 1 {
        fmt.Println("")
      }
    }
  }
}
