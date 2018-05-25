package main

import (
	"bufio"
	"fmt"
	"math"
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

	var n, k int
	scanf("%d %d\n", &n, &k)
	array := make([]int, n, int(2*math.Pow(10, 5)))

	for i := 0; i < n; i++ {
		scanf("%d", &array[i])
	}

	sort.Ints(array)

	answer := 0

	// printf("sorted array: %v\n", array)

	if k == 0 {
		answer = array[0] - 1
	} else {
		answer = array[k-1]
	}

	// printf("answer: %v\n", answer)

	counter := 0

	for _, val := range array {
		if val <= answer {
			counter++
		}
	}

	if answer < 1 || counter != k {
		printf("-1\n")
	} else {
		printf("%v\n", answer)
	}

	// example
	// var a, b int
	// scanf("%d %d", &a, &b)
	// printf("%d\n", a+b)
}
