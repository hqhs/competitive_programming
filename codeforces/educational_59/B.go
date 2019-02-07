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

func dr(n int) int {
	return 1 + (n - 1) % 9
}

func sol(k int64, x int64) int64 {
	return x + (k - 1) * 9
}

func main() {
	//STDOUT MUST BE FLUSHED MANUALLY
	defer writer.Flush()

	var n int
	var x, k int64
	scanf("%d\n", &n)
	for i := 0; i < n; i++ {
		scanf("%d %d\n", &k, &x)
		printf("%d\n", sol(k, x))
		// for f := 1; f <= 10; f++ {
		// 	printf("%d: %d\t", i*10 + f, dr(i*10 + f))
		// }
		// printf("\n")
	}
}
