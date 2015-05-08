package main
import "fmt"
import "math"

func main() {
  var t int
  var s string
  fmt.Scan(&t)

  for i := 0; i < t; i++ {
    fmt.Scan(&s)
    s_len := len(s)
    is_funny_string := true
    for j := 1; j < s_len; j++ {
      abs_s := math.Abs(float64(s[j]) - float64(s[j - 1]))
      abs_r := math.Abs(float64(s[s_len - j - 1]) - float64(s[s_len - j]))

      if abs_s != abs_r {
        is_funny_string = false
        break
      }
    }

    if is_funny_string {
      fmt.Println("Funny")
    } else {
      fmt.Println("Not Funny")
    }
  }
}

