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
	var input string
	// int('a') == 97, int('z') == 122
	scanf("%s\n", &input)
	n := len(input)
	const chunkSize = 3
	// small hack for lesser memory allocations
	tmp := make([]byte, chunkSize*n)
	var triplets [][]byte
	for i := 0; i < len(tmp); i += chunkSize {
		end := i + chunkSize
		triplets = append(triplets, tmp[i:end])
	}
	// now we need to generate unique triplets from 0 to n
	// O(n) solution
	req1, req2, req3 := make([]byte, n), make([]byte, n), make([]byte, n)
	for i := 0; i < n; i++ {
		req1[i] = byte('a' + i%26)
		req2[i] = byte('a' + (i/26)%26)
		req3[i] = byte('a' + ((i/26)/26)%26)
	}
	// printf("triplets: %+v\n", triplets)
	var ans1, ans2, ans3 string
	printf("? %v\n", string(req1))
	writer.Flush()
	scanf("%s\n", &ans1)
	printf("? %v\n", string(req2))
	writer.Flush()
	scanf("%s\n", &ans2)
	printf("? %v\n", string(req3))
	writer.Flush()
	scanf("%s\n", &ans3)
	// so we need to find (i, j) pairs with brute force
	swaps := make([]int, n)
	for i := 0; i < n; i++ {
		swaps[i] = (int(ans1[i]) - 'a') + (int(ans2[i])-'a')*26 + (int(ans3[i])-'a')*26*26
		// printf("swaps[%d] = %v\n", i, swaps[i])
	}
	// printf("sorted swaps: %+v\n", swaps)
	original := make([]byte, n)
	for i := range original {
		original[i] = 'a'
	}
	for i := 0; i < n; i++ {
		// original[i] = input[swaps[i]]
		original[swaps[i]] = input[i]
	}
	printf("! %v\n", string(original))
	writer.Flush()
}

type swap struct {
	i int
	j int
}
