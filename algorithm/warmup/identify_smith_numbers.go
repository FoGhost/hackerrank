package main
import "fmt"

func main() {
  var n int
  fmt.Scan(&n)
  raw_n := n

  factors := []int{}
  for n % 2 == 0 {
    n = n / 2
    factors = append(factors, 2)
  }

  for i := 3; i <= n; {
    if n % i == 0 {
      n = n / i
      factors = append(factors, i)
    } else {
      i += 2
    }
  }

  // if n is not a composite number
  if len(factors) == 1 {
    fmt.Println(0)
    return
  }

  sum_factors_digits := 0
  for _, num := range factors {
    sum_factors_digits += sumDigits(num)
  }

  if sum_factors_digits == sumDigits(raw_n) {
    // n is a Smith number
    fmt.Println(1)
  } else {
    // n is not a Smith number
    fmt.Println(0)
  }
}

func sumDigits(num int) int {
  sum := 0
  for num > 0 {
    sum += num % 10
    num = num / 10
  }

  return sum
}
