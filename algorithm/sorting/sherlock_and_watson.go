package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)

	var n, k, q int
	fmt.Fscan(reader, &n, &k, &q)

	ar := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &ar[i])
	}

	k = k % n
	new_ar := ar[n-k : n]
	new_ar = append(new_ar, ar[0:n-k]...)
	var q_index int
	for j := 0; j < q; j++ {
		fmt.Fscan(reader, &q_index)
		fmt.Fprintln(writer, new_ar[q_index])
	}
	writer.Flush()
}
