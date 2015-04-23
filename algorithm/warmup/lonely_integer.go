package main
import "fmt"
import "sort"

func main() {
  var n int
  fmt.Scan(&n)

  var a = make([]int, n)
  for i := 0; i < n; i++ {
    fmt.Scan(&a[i])
  }
  sort.Ints(a)
  for i := 0; i < n; i++ {
    if i == 0 {
      if i == n -1 {
        fmt.Println(a[i])
      } else if a[i] != a[i+1] {
        fmt.Println(a[i])
      }
    } else if i == (n - 1) {
      if (a[i] != a[i-1]) {
        fmt.Println(a[i])
      }
    } else {
      if (a[i] != a[i+1] && a[i] != a[i-1]) {
        fmt.Println(a[i])
      }
    }
  }
}
