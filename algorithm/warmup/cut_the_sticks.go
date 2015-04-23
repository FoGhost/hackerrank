package main
import "fmt"
import "sort"

func main() {
  var n int
  fmt.Scan(&n)

  var stick_arr = make([]int, n)
  for i := 0; i < n; i++ {
    fmt.Scan(&stick_arr[i])
  }

  sort.Ints(stick_arr)
  stick_num := n

  min := stick_arr[0]
  new_min := stick_arr[0]
  for stick_num > 0 {
    fmt.Println(stick_num)

    for i := n - 1; i >= 0; i-- {
      if stick_arr[i] == 0 { break }
      stick_arr[i] = stick_arr[i] - min
      if stick_arr[i] == 0 {
        if (i + 1) < n && stick_arr[i + 1] > 0 {
          new_min = stick_arr[i + 1] // tmp save new min
        }
        //Fix update min
        stick_num = stick_num - 1 // update stick+num
      }
    }
    min = new_min // update min
  }
}
