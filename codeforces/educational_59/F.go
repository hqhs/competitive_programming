package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func min(i, k int) int64 {
	if i < k {
		return int64(i)
	}
	return int64(k)
}

func main() {
	//STDOUT MUST BE FLUSHED MANUALLY
	defer writer.Flush()

	var n int

	scanf("%d\n", &n)
	up := make([]int, n+5)
	down := make([]int, n+5)
	k := make([]int, n+5)

	for i := 0; i < n; i++ {
		scanf("%d %d %d\n", &up[i], &down[i], &k[i])
	}

	mat := make([][]int64, n)
	for i := range mat {
		mat[i] = make([]int64, n)
	}

	for j := 0; j < n; j++ {
		for i := 0; i < n; i++ {
			mat[j][i] = int64(up[j]) - min(i, k[j])*int64(down[j])
		}
	}

	for i := 0; i < n; i++ {
		printf("%v\n", mat[i])
	}
	// TODO: implement this:
	// https://www.topcoder.com/community/competitive-programming/tutorials/assignment-problem-and-hungarian-algorithm/
}
