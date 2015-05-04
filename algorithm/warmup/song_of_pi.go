package main
import (
  "fmt"
  "bufio"
  "os"
)

func main() {
  var t int
  fmt.Scan(&t)

  bio := bufio.NewReader(os.Stdin)
  pi_digits := [29]int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5, 8, 9, 7, 9, 3, 2, 3, 8, 4, 6, 2, 6, 4, 3, 3, 8, 3, 3}

  for i := 0; i < t; i++ {
    line, _ := bio.ReadBytes('\n')
    char_count := 0
    word_index := 0
    is_pi_song := true
    for _, char := range line {
      if ((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')) {
        char_count += 1
      } else {
        if char_count == pi_digits[word_index] {
          word_index += 1
          // reset char_count
          char_count = 0
        } else {
          // It's not a pi song
          is_pi_song = false
          break
        }
      }
    }

    if is_pi_song {
      fmt.Println("It's a pi song.")
    } else {
      fmt.Println("It's not a pi song.")
    }
  }
}
