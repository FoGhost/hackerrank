package main
import "fmt"

func main() {
  var t int
  fmt.Scan(&t)

  for i := 0; i < t; i++ {
    var b, w int
    fmt.Scan(&b, &w)
    var x, y, z int
    fmt.Scan(&x, &y, &z)

    if y < x && (y + z) < x {
      fmt.Println((y + z) * b + y * w)
    } else if x < y && (x + z) < y {
      fmt.Println(x * b + (x + z) * w)
    } else {
      fmt.Println(x * b + y * w)
    }
  }
}
