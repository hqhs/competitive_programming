package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func main() {
	//STDOUT MUST BE FLUSHED MANUALLY
	defer writer.Flush()
	var n int
	scanf("%d", &n)
	in := make([]int, n)
	for i := range in {
		scanf("%d\n", &in[i])
	}
	// printf("%v\n")
}
