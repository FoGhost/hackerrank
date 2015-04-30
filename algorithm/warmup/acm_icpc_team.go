package main
import "fmt"

func countTopics(p_a string, p_b string) int {
  count := 0
  for i := 0; i < len(p_a); i++ {
    if p_a[i] == '1' || p_b[i] == '1' {
      count = count + 1
    }
  }

  return count
}

func main() {
  var n, m int
  fmt.Scan(&n, &m)

  people := make([]string, n)
  for i := 0; i < n; i++ {
    var p_topics string
    fmt.Scan(&p_topics)
    people[i] = p_topics
  }

  team_topics := make(map[int]int)
  for i := 0; i < n; i++ {
    for j := i + 1; j < n; j++ {
      topics_count := countTopics(people[i], people[j])
      if val, present := team_topics[topics_count]; present {
        team_topics[topics_count] = val + 1
      } else {
        team_topics[topics_count] = 1
      }
    }
  }

  max_topics_count := 0
  for topics_count, _ := range team_topics {
    if topics_count > max_topics_count {
      max_topics_count = topics_count
    }
  }
  fmt.Println(max_topics_count)
  fmt.Println(team_topics[max_topics_count])
}
