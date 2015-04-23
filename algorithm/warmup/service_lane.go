package main
import "fmt"

func main() {
  var n, t int
  fmt.Scanln(&n, &t)
  width := make([]int, n)
  for i := range width {
    var w int
    fmt.Scan(&w)
    width[i] = w
  }

  for i := 0 ; i < t; i++ {
    var in, out int
    vehicle := 3
    fmt.Scan(&in, &out)
    for j := in ; j <= out ; j++ {
      w := width[j]
      if (vehicle > w) {
        vehicle = w
      }
    }
    fmt.Println(vehicle)
  }
}
