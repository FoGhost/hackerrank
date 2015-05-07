package main
import "fmt"
import "strconv"

func numberToWords(num int) string {
  if num > 59 {
    return "unkown"
  }

  number_below_twenty_map := []string{
    "zero",
    "one",
    "two",
    "three",
    "four",
    "five",
    "six",
    "seven",
    "eight",
    "nine",
    "ten",
    "eleven",
    "twelve",
    "thirteen",
    "fourteen",
    "fifteen",
    "sixteen",
    "seventeen",
    "eighteen",
    "nineteen",
  }
  number_above_twenty_map := []string{
    "twenty",
    "thirty",
    "forty",
    "fifty",
  }

  var words string
  if num < 20 {
    words =  number_below_twenty_map[num]
  } else {
    i := num / 10
    j := num % 10

    if j == 0 {
      words = fmt.Sprintf("%s", number_above_twenty_map[i - 2])
    } else {
      words = fmt.Sprintf("%s %s", number_above_twenty_map[i - 2], number_below_twenty_map[j])
    }
  }

  return words
}

func timeToWords(h int, m int) string {
  var words string
  switch {
    case m == 0:
      words = fmt.Sprintf("%s o' clock", numberToWords(h))
    case m == 1:
      words = fmt.Sprintf("one minute past %s", numberToWords(h))
    case m > 1 && m < 30:
      if m == 15 {
        words = fmt.Sprintf("quarter past %s", numberToWords(h))
      } else {
        words = fmt.Sprintf("%s minutes past %s", numberToWords(m), numberToWords(h))
      }
    case m == 30:
      words = fmt.Sprintf("half past %s", numberToWords(h))
    case m > 30 && m <= 59:
      if m == 45 {
        words = fmt.Sprintf("quarter to %s", numberToWords(h + 1))
      } else {
        words = fmt.Sprintf("%s minutes to %s", numberToWords(60 - m), numberToWords(h + 1))
      }
  }

  return words
}

func main() {
  var h, m string
  fmt.Scan(&h, &m)
  m_int, _ := strconv.Atoi(m)
  h_int, _ := strconv.Atoi(h)
  fmt.Println(timeToWords(h_int, m_int))
}
