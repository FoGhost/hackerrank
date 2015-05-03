package main
import "fmt"
import "sort"

type intArray []int
func (s intArray) Len() int { return len(s) }
func (s intArray) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s intArray) Less(i, j int) bool { return s[i] < s[j] }

func main() {
  var n, k int
  fmt.Scan(&n, &k)

  arr := make([]int, n)
  for i := 0; i < n; i++ {
    fmt.Scan(&arr[i])
  }

  sort.Sort(intArray(arr))

  var min, max, unfairness int
  for j := 0; j < n - (k - 1); j++ {
    min = arr[j]
    max = arr[j + k - 1]

    if j == 0 {
      unfairness = max - min
    } else {
      if (max - min) < unfairness {
        unfairness = max - min
      }
    }
  }
  fmt.Println(unfairness)
}
