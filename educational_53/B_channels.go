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

func pubChan(s []int, c chan int) {
	for _, val := range s {
		c <- val
	}
	close(c)
}

func solution(stack, sequence []int) []int {
	solSeq := make([]int, len(sequence))
	had := make(map[int]bool)

	s, q := make(chan int, len(sequence)), make(chan int, len(sequence))
	pubChan(stack, s)
	pubChan(sequence, q)

	index, counter := 0, 0

	for qVal := range q {
		if _, ok := had[qVal]; ok {
			// printf("skipping value: %v\n", qVal)
			index++
			continue
		}
		counter = 0

		for sVal := range s {
			// printf("inside inner loop:%v\n", sVal)
			counter++
			if sVal == qVal {
				solSeq[index] = counter
				index++
				// printf("breaked inner loop: %v\n", sVal)
				break
			}
			had[sVal] = true
		}
	}

	return solSeq
}

func main() {
	//STDOUT MUST BE FLUSHED MANUALLY
	defer writer.Flush()

	var n int
	scanf("%d\n", &n)
	stack, sequence := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		scanf("%d", &stack[i])
	}
	scanf("\n")
	for i := 0; i < n; i++ {
		scanf("%d", &sequence[i])
	}
	scanf("\n")
	for _, val := range solution(stack, sequence) {
		printf("%v ", val)
	}
	printf("\n")
}
