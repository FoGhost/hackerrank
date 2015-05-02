package main
import "fmt"

func gcd(a int, b int) int {
  if a < b {
    a, b = b, a
  }

  c := a % b
  if c == 0 {
    return b
  } else {
    return gcd(b, c)
  }
}

func isExistSubset(arr_a []int) bool{
  for i, a := range arr_a {
    for j := i + 1; j < len(arr_a); j++ {
      if a == arr_a[j] {
        continue
      }

      found_gcd := gcd(a, arr_a[j])
      if found_gcd == 1 {
        return true
      }
    }
  }
  return false
}

func main() {
  var t int
  fmt.Scan(&t)

  for i := 0; i < t; i++ {
    var n int
    fmt.Scan(&n)

    arr_a := []int{}
    for j := 0; j < n; j++ {
      var a int
      fmt.Scan(&a)
      arr_a = append(arr_a, a)
    }

    if isExistSubset(arr_a) {
      fmt.Println("YES")
    } else {
      fmt.Println("NO")
    }
  }
}

