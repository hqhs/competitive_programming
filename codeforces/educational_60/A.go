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

	var n, max, tmp, counter, maxCounter int
	max = -1
	scanf("%d\n", &n)
	// input := make([]int, n)
	for i := 0; i < n; i++ {
		scanf("%d ", &tmp)
		if tmp > max {
			max = tmp
			counter = 1
			maxCounter = 1
		} else if tmp == max {
			counter++
		} else if tmp != max {
			counter = 0
		}
		if counter > maxCounter {
			maxCounter = counter
		}
	}
	// count maximum length of sequence where each elem equal to max
	printf("%+v\n", maxCounter)
}
