package main

// test input:
// 5 21
// 2 4 100 2 6
// output
// 6

// test input:
// 3 38
// 5 2 5
// ouput
// 10

import (
	"bufio"
	"fmt"
	_ "math"
	"os"
	_ "sort"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func sum(slice []int) uint64 {
	sum := uint64(0)
	for _, value := range slice {
		sum += uint64(value)
	}
	return sum
}

// return index INCLUDED
func binarySearch(targetSlic []int, T uint64) int {

	startIndex := 0
	endIndex := len(targetSlic) - 1

	for startIndex <= endIndex {
		// fmt.Printf("%v, %v\n", startIndex, endIndex)

		median := startIndex + (endIndex-startIndex)/2

		// fmt.Printf("sum: %v\n", sum(targetSlic[0:median]))
		if uint64(sum(targetSlic[0:median])) <= T {
			// fmt.Print("less")
			startIndex = median + 1
		} else {
			// fmt.Print("more")
			endIndex = median - 1
		}
	}

	return startIndex - 1
}

func solution(prices []int, T uint64) uint64 {
	min := prices[0]

	for _, val := range prices {
		// printf("min: %v\n", min)
		if val < min {
			min = val
		}
	}

	candies := uint64(0)

	s := uint64(sum(prices))
	// find s = sum(prices), if s < T, optimize, if not, exlude max, check s < T, repeat until s < T
	// break then min > T
	for s != 0 {
		// skip steps, if sum less then current balance
		if s <= T {
			candies += uint64(len(prices)) * (T / uint64(s))
			T %= uint64(s)
		} else {
			tmpS, tmpT, tmpCandies := uint64(0), uint64(T), 0

			pricesPtr := prices[:0]
			for _, val := range prices {
				// could we buy candy?
				if uint64(val) <= tmpT {
					tmpS += uint64(val)
					tmpT -= uint64(val)
					tmpCandies++
					// we bought it, add to prices
					pricesPtr = append(pricesPtr, val)
				}
			}
			// printf("to exlude: %v\n", pricesPtr)
			candies += uint64(tmpCandies)
			prices = pricesPtr
			s = tmpS
			if tmpS <= T {
				T -= tmpS
			} else {
				break
			}
		}
	}

	return uint64(candies)
}

func main() {
	//STDOUT MUST BE FLUSHED MANUALLY
	defer writer.Flush()

	// printf("maxInt: %v\n", math.MaxUint64)
	var n int
	var T uint64
	scanf("%d %d", &n, &T)
	prices := make([]int, n)
	scanf("\n")
	for i := 0; i < n; i++ {
		scanf("%d", &prices[i])
	}
	printf("%v\n", solution(prices, T))
}
