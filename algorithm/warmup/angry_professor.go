package main
import "fmt"

func willBeCancelled(n int, k int, students []int) bool {
  arrived_count := 0
  for i := 0; i < n; i++ {
    if students[i] <= 0 {
      arrived_count = arrived_count + 1
    }
    if arrived_count >= k {
      return false
    }
  }
  return true
}


func main() {
  var case_num int
  fmt.Scan(&case_num)

  for i := 0; i < case_num; i++ {
    var n, k int
    fmt.Scan(&n, &k)
    students := make([]int, n)
    for j := 0; j < n; j++ {
      fmt.Scan(&students[j])
    }
    ret := willBeCancelled(n, k, students)
    if ret {
      fmt.Println("YES")
    } else {
      fmt.Println("NO")
    }
  }
}

