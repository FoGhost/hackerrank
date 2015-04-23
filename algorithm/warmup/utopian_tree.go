package main
import "fmt"

func main() {
  var n int
  var cycles_num int
  fmt.Scan(&n)

  for i := 0; i < n; i++ {
    var height uint = 1

    fmt.Scan(&cycles_num)
    for j :=1; j <= cycles_num ; j++ {
      if j % 2 == 0 {
        height = height + 1
      } else {
        height = height * 2
      }
    }
    fmt.Println(height)
  }
}
