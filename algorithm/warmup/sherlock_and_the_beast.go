package main
import "fmt"

func printLargestNum(digit_5_count int, digit_3_count int) {
  string_buffer := make([]byte, digit_5_count + digit_3_count)
  buffer_size := 0
  for i := 0; i < digit_5_count; i++ {
    buffer_size += copy(string_buffer[buffer_size:], "5")
  }
  for i := 0; i < digit_3_count; i++ {
    buffer_size += copy(string_buffer[buffer_size:], "3")
  }
  fmt.Println(string(string_buffer))
}

func main() {
  var t int
  fmt.Scan(&t)

  for i := 0; i < t; i++ {
    var n int
    fmt.Scan(&n)

    digit_5_count := 0
    digit_3_count := 0
    found_largest_num := false
    for n >= 0 {
      if n % 3 == 0 {
        digit_5_count = n
        found_largest_num = true
        break
      } else {
        digit_3_count = digit_3_count + 5
        n = n - 5
      }
    }

    if found_largest_num {
      printLargestNum(digit_5_count, digit_3_count)
    } else {
      fmt.Println(-1)
    }
  }
}
