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

func main() {
	//STDOUT MUST BE FLUSHED MANUALLY
	defer writer.Flush()

	//count how many elements before i less then a[i] and find max(a[i])

	var a, b int
	scanf("%d %d", &a, &b)
	printf("%d\n", a+b)
}
