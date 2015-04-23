package main
import "fmt"

func main() {
  var l, r int
  var max int = 0
  var xor_res int = 0
  fmt.Scan(&l, &r)

  for i := l; i <= r; i++ {
    for j :=l; j <= r; j++ {
      xor_res = i ^ j
      if max < xor_res {
        max = xor_res
      }
    }
  }
  fmt.Println(max)
}
