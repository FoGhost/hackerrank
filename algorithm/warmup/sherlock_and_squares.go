package main
import "fmt"
import "math"

func main() {
  var t int
  fmt.Scan(&t)

  for i := 0; i < t; i++ {
    var a, b int;
    fmt.Scan(&a, &b)

    square_count := 0
    square_root_a := int(math.Sqrt(float64(a)))
    if (square_root_a * square_root_a == a) {
      square_count = square_count + 1
    }
    square_root_a = square_root_a + 1
    next_a := square_root_a * square_root_a
    for next_a <= b {
      square_count = square_count + 1
      square_root_a = square_root_a + 1
      next_a = square_root_a * square_root_a
    }
    fmt.Println(square_count)
  }
}
