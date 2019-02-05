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

type ex struct {
	pages int
	start int
	end   int
	step  int
}

func (e ex) String() string {
	return fmt.Sprintf("Pages: %d, start: %d, end: %d, step: %d", e.pages, e.start, e.end, e.step)
}

func (e ex) Eval() int {
	buf := make([]bool, e.pages)
	i := e.start
	counter := 0
	for {
		if buf[i] {
			return -1
		} else if i == e.end {
			return counter
		} else {

		}
	}
}

func main() {
	//STDOUT MUST BE FLUSHED MANUALLY
	defer writer.Flush()

	var n int
	scanf("%d\n", &n)
	data := make([]ex, n)
	for i := range data {
		scanf("%d %d %d %d\n", &data[i].pages, &data[i].start, &data[i].end, &data[i].step)
	}

	for _, val := range data {
		fmt.Println(val.Eval())
	}

	// var a, b int
	// scanf("%d %d", &a, &b)
	// printf("%d\n", a+b)
}
