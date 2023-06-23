package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()

	var n int

	fmt.Fscan(r, &n)

	datas := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(r, &datas[i])
	}

	var sum, result int

	sort.Ints(datas)

	for i := 0; i < n; i++ {
		var startIdx int
		endIdx := n - 1
		find := datas[i]

		for startIdx < endIdx {
			sum = datas[startIdx] + datas[endIdx]
			if sum == find {
				if startIdx != i && endIdx != i {
					result++
					break
				} else if startIdx == i {
					startIdx++
				} else if endIdx == i {
					endIdx--
				}
			} else if sum < find {
				startIdx++
			} else {
				endIdx--
			}
		}
	}

	fmt.Fprint(w, result)
}