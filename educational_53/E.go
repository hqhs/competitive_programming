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

func intervalSum(
	n uint64, // n is end of interval, k is allowed amount of numbers
	k int) uint64 {

	return uint64(0)
}

func main() {
	//STDOUT MUST BE FLUSHED MANUALLY
	defer writer.Flush()
	var l, r, k uint64
	scanf("%d %d %d", &l, &r, &k)
	answer := intevalSum(r) - intervalSum(l-1)
	printf("%d\n", answer)
}
