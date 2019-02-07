package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func main() {
	//STDOUT MUST BE FLUSHED MANUALLY
	defer writer.Flush()

	var n, len, num int
	var d string
	scanf("%d\n", &n)
	for i := 0; i < n; i++ {
		scanf("%d\n", &len)
		scanf("%s\n", &d)
		// printf("len: %v, data: %v\n", len, d)
		if len > 2 {
			printf("YES\n")
			printf("2\n")
			printf("%v %v\n", d[:1], d[1:])
			continue
		}
		num, _ = strconv.Atoi(d)
		if num / 10 >= num % 10 {
			printf("NO\n")
			continue
		}
		printf("YES\n")
		printf("2\n")
		printf("%v %v\n", d[:1], d[1:])
	}
	// printf("%d\n", a+b)
}
