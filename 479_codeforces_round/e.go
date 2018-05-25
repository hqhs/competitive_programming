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

func dfs() bool { // true if returned to start point

}

func findNextFalse(flags *[]bool) func() int {
	return func() int {
		for i, val := range *flags {
			if !val {
				(*flags)[i] = true
				return i
			}
		}
		return -1
	}
}

func main() {
	//STDOUT MUST BE FLUSHED MANUALLY
	defer writer.Flush()

	var n, m, place, ptr int
	scanf("%d %d\n", &n, &m)

	ptrs := make(map[int][]int, n+1)

	flags := make([]bool, n+1) // default is false

	for i := 0; i < m; i++ {
		scanf("%d %d\n", &place, &ptr)
		ptrs[place] = append(ptrs[place], ptr)
		ptrs[ptr] = append(ptrs[ptr], place)
	}

	flags[0] = true // becouse it's easier to start with 1 :D
	f := findNextFalse(&flags)
	loops_counter := 0

	for i := 0; i != -1; i = f() { // iterate over flags
		printf("i: %v\n", i)
		if dfs() && len(ptrs[i]) <= 2 {
			loops_counter++
		}
	}

	printf("%v\n", loops_counter)

	// printf("ptrs: %v\n", ptrs)

	// var a, b int
	// scanf("%d %d", &a, &b)
	// printf("%d\n", a+b)
}
