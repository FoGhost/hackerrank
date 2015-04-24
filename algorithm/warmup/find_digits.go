package main
import "fmt"

func findDigits(n int64) int {
  divided_n := n
  count := 0
  for {
    remainer := divided_n % 10
    if (remainer != 0 && n % remainer == 0) {
      count = count + 1
    }

    divided_n = divided_n / 10
    if divided_n == 0 {
      break
    }
  }

  return count
}

func main() {
  var t int
  fmt.Scan(&t)

  for i := 0; i < t; i++ {
    var n int64
    fmt.Scan(&n)
    ret := findDigits(n)
    fmt.Println(ret)
  }
}
