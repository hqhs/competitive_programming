package main

import (
	"unicode"
	"bufio"
	"fmt"
	"os"
	"math/big"
	"sort"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func main() {
	//STDOUT MUST BE FLUSHED MANUALLY
	defer writer.Flush()

	max := big.NewInt(0)

	var n, k int
	scanf("%d %d\n", &n, &k)
	dmg := make([]int64, n)
	for i := 0; i < n; i++ {
		scanf("%d ", &dmg[i])
	}
	seq := make([]rune, n)

	var t rune
	for i := 0; i < n;  {
		scanf("%c", &t)
		if !unicode.IsSpace(t) {
			seq[i] = t
			i++
		}
	}
	// SOLUTION
	var lp, rp int
	i := 0
	for {
		if i >= len(seq) - 1 {
			// special case: len(seq) == 1
			if len(seq) == 1 {
				max.Add(max, big.NewInt(dmg[0]))
				break
			}
			// add last one
			if i == len(seq) - 1 && seq[i-1] != seq[i] {
				max.Add(max, big.NewInt(dmg[i]))
			}
			break
		}
		if seq[i] != seq[i+1] {
			// printf("add: %d", dmg[i])
			max.Add(max, big.NewInt(dmg[i]))
			// printf("\tmax: %v\n", max)
			i++
			continue
		}
		// find sequence with same chars
		lp = i
		rp = i
		for {
			rp++
			if seq[lp] != seq[rp] {
				rp--
				break
			} else if rp == len(seq) - 1 {
				break
			}
		}
		// printf("same chars: %+v\n", dmg[lp:rp+1])
		tmp := make([]int64, len(dmg[lp:rp+1]))
		copy(tmp, dmg[lp:rp+1])
		sort.Slice(tmp, func(i, j int) bool { return tmp[i] < tmp[j] })
		// printf("Sorted tmp: %+v\n", tmp)
		// get k max elements from tmp
		for i := 0; i < k && i < len(tmp); i++  {
			curMax := tmp[len(tmp)-1-i]
			// printf("add: %d", curMax)
			max.Add(max, big.NewInt(curMax))
			// printf("\tmax: %v\n", max)
		}
		// printf("max: %v\n", max)
		i = rp + 1
	}
	printf("%v\n", max)
}
