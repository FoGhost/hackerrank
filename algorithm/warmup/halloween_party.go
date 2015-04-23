package main
import "fmt"

func main() {
  var t int
  fmt.Scan(&t)

  for i := 0; i < t; i++ {
    var k int
    fmt.Scan(&k)
    if k % 2 == 0 {
      cut_h := k / 2
      cut_v := k / 2
      fmt.Println(cut_h * cut_v)
    } else {
      cut_h := k / 2
      cut_v := cut_h + 1
      fmt.Println(cut_h * cut_v)
    }
  }
}
