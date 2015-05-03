package main
import (
  "fmt"
  "math"
)

func main() {
  var n, m int64
  var a, b, k int64
  fmt.Scan(&n, &m)

  var total_candy_num int64
  for i := int64(0); i < m; i++ {
    fmt.Scan(&a, &b, &k)
    total_candy_num += (b - a + 1) * k
  }
  fmt.Println(int64(math.Floor(float64(total_candy_num / n))))
}
