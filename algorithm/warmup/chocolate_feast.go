package main
import "fmt"

func main() {
  var t int
  fmt.Scan(&t)

  for i := 0; i < t; i++ {
    var n, c, m int
    fmt.Scan(&n, &c, &m)

    chocolate_num := n / c
    wrappers_num := chocolate_num
    for wrappers_num >= m {
      new_chocolate_num := wrappers_num / m
      chocolate_num = chocolate_num + new_chocolate_num
      wrappers_num = new_chocolate_num + (wrappers_num % m)
    }

    fmt.Println(chocolate_num)
  }
}
