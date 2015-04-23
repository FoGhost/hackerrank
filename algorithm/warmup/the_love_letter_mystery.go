package main
import "fmt"

func main() {
  var t int
  fmt.Scan(&t)

  for i := 0 ; i < t ; i++ {
    var word string
    fmt.Scan(&word)
    l := len(word) // 2
    count := 0
    for i := 0; i < (l / 2); i++ {
      left := int(word[i])
      right := int(word[(l-1) - i])
      if left > right {
        count = count + left - right
      } else {
        count = count + right - left
      }
    }
    fmt.Println(count)
  }
}
