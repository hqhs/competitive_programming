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

func dfs(matrix [][]uint64, used []bool, i int, n uint64) []uint64 {
	comp := make([]uint64, 0, n)

	for i := range matrix[i] {
		if !used[i] {

		}
	}

	return comp
}

func main() {
	//STDOUT MUST BE FLUSHED MANUALLY
	defer writer.Flush()

	var n, m, place, ptr uint64
	scanf("%d %d\n", &n, &m)

	matrix := make([][]uint64, n)
	degrees := make([]uint64, n) // default is 0
	for i := uint64(0); i < m; i++ {
		scanf("%d %d\n", &place, &ptr)
		place--
		ptr--
		matrix[place] = append(matrix[place], ptr)
		matrix[ptr] = append(matrix[ptr], place)
		degrees[place]++
		degrees[ptr]++
	}

	fmt.Println(matrix)
	fmt.Println(degrees)

	flags := make([]bool, n+1) // default is false

	for i, value := range flags {
		if !value {
			// var ok bool

			comp := dfs(used[:], i, n)
			fmt.Println(comp)
		}
	}

	// loopsCounter := 0

	// printf("%v\n", loops_counter)
}
