package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

type delimeters struct {
	twos   int
	threes int
}

type sortDelimeters []delimeters

func (arr sortDelimeters) Len() int {
	return len(arr)
}

func (arr sortDelimeters) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func (arr sortDelimeters) Less(i, j int) bool {
	return arr[i].twos <= arr[j].twos && arr[i].threes >= arr[j].threes
}

func getDelimeters(n uint64) delimeters {
	r := delimeters{0, 0}
	for {
		if n%2 == 0 {
			r.twos++
			n /= 2
		} else if n%3 == 0 {
			r.threes++
			n /= 3
		} else {
			break
		}
	}
	return r
}

func insert(num uint64, i int, array *[]uint64) {
	*array = append(*array, 0)
	buffer := (*array)[i]
	(*array)[i] = num
	// printf("buffer: %v, a[i]: %v\n", buffer, (*array)[i])
	for f := i + 1; f < len((*array)); f++ {
		printf("f == %v\n", f)
		buffer, (*array)[f] = (*array)[f], buffer
	}
	// printf("array: %v\n", *array)
}

func main() {
	//STDOUT MUST BE FLUSHED MANUALLY
	defer writer.Flush()

	var n int
	scanf("%d\n", &n)
	arr := make([]uint64, n)
	sortedArr := make([]delimeters, n)
	answerArr := make([]uint64, n)
	delCounter := make(map[delimeters]uint64)

	for i := range arr {
		scanf("%d", &arr[i])
		del := getDelimeters(arr[i])
		delCounter[del] = arr[i]
		sortedArr[i] = del
	}

	// printf("sortedArr: %v\n", sortedArr)

	sort.Sort(sortDelimeters(sortedArr))

	// printf("sortedArr: %v\n", sortedArr)

	for i, val := range sortedArr {
		answerArr[i] = delCounter[val]
	}

	for _, val := range answerArr {
		printf("%v ", val)
	}

	printf("\n")

}
