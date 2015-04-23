package main
import "fmt"

func main() {
  var n int
  var a, b uint
  fmt.Scan(&n)

  for i := 0; i < n; i++ {
    var sum uint
    fmt.Scan(&a, &b)
    sum = a + b
    fmt.Println(sum)
  }
}
