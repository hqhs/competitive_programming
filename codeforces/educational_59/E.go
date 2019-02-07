package main

import (
	"bufio"
	"fmt"
	"os"
	"math"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

type dpcall struct {
	c   byte
	l   int
	r   int
	cnt int
}

var dp map[dpcall]*int64

type anscall struct {
	l int
	r int
}

var ans map[anscall]*int64

var s []byte
var costs []int

func main() {
	//STDOUT MUST BE FLUSHED MANUALLY
	defer writer.Flush()

	var n int
	scanf("%d\n", &n)
	s = make([]byte, n)
	var tmp string
	scanf("%s\n", &tmp)
	for i, c := range tmp {
		s[i] = byte(c) - byte('0')
	}

	costs = make([]int, n)
	for i := 0; i < n; i++ {
		scanf("%d ", &costs[i])
	}
	// printf("scanned: %v\n", costs)

	// SOLUTION
	// init memoization
	ans = make(map[anscall]*int64)
	dp = make(map[dpcall]*int64)
	printf("%v\n", calcAns(0, n))
}

func calcAns(l, r int) int64 {
	if l >= r {
		return 0
	}
	cal, ok := ans[anscall{l, r}]
	if ok {
		return *cal
	}
	res := int64(0)
	ans[anscall{l, r}] = &res

	for cnt := 1; cnt <= r - l; cnt++ {
		res = max(res, calcDp(0, l, r, cnt) + int64(costs[cnt - 1]))
		res = max(res, calcDp(1, l, r, cnt) + int64(costs[cnt - 1]))
		// printf("res for l(%d), r(%d): %d\n", l, r, res)
	}
	// printf("calcAns(%+v) : %d (cached: %d)\n", anscall{l, r}, res, *ans[anscall{l, r}])
	return res
}

func calcDp(c byte, l, r, cnt int) int64 {
	if cnt == 0 {
		return calcAns(l, r)
	}
	val, ok := dp[dpcall{c, l, r, cnt}]
	if ok {
		return *val
	}
	res := -int64(math.MaxInt64)
	dp[dpcall{c, l, r, cnt}] = &res

	for i := l; i < r; i++ {
		if (c == s[i]) {
			res = max(res, calcAns(l, i) + calcDp(c, i + 1, r, cnt - 1))
		}
	}
	// printf("calcDp(%+v) : %d (cached: %d)\n", dpcall{c, l, r, cnt}, res, *dp[dpcall{c, l, r, cnt}])
	return res
}

func max(i, f int64) int64 {
	if i > f {
		return i
	}
	return f
}
