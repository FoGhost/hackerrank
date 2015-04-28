package main
import "fmt"

func main() {
  var t int
  fmt.Scan(&t)

  for i := 0; i < t; i++ {
    var n, a, b int
    fmt.Scan(&n, &a, &b)

    if (a == b) {
      fmt.Println((n-1) * a)
    } else {
      var diff, min, max int
      if a > b {
        diff = a - b
        min = b * (n - 1)
        max = a * (n - 1)
      } else {
        diff = b - a
        min = a * (n - 1)
        max = b * (n - 1)
      }

      j := min
      fmt.Printf("%d", j)
      for j < max {
        j = j + diff
        fmt.Printf("% d", j)
      }
      fmt.Println("")
    }

  }
}
