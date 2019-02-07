package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	_ "reflect"
)

// input optimization
var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

type pos struct {
	x int
	y int
}

// globals

var end pos
var s string

// solution
func abs(i int) int {
	return int(math.Abs(float64(i)))
}

func (cur pos) String() string {
	return fmt.Sprintf("x: %v, y: %v\n", cur.x, cur.y)
}

func (cur *pos) update(mv rune, d int) {
	switch mv {
	case 'U':
		cur.y += d
	case 'D':
		cur.y -= d
	case 'L':
		cur.x -= d
	case 'R':
		cur.x += d
	}
}

func can(start, end pos, l int) bool {
	d := abs(start.x-end.x) + abs(start.y-end.y)
	if d%2 != l%2 {
		return false
	}
	return l >= d
}

func ok(length int) bool {
	p := pos{0, 0}
	for _, value := range s[length:] {
		p.update(value, 1)
	}

	l, r := 0, length

	for {
		if can(p, end, length) {
			return true
		}

		if r == len(s) {
			break
		}
		p.update(rune(s[l]), 1)
		l++
		p.update(rune(s[r]), -1)
		r++
	}

	return false
}

func solution() int {
	if !ok(len(s)) {
		return -1
	}

	l, r := -1, len(s)
	for r-l > 1 {
		mid := (l + r) / 2
		if ok(mid) {
			r = mid
		} else {
			l = mid
		}
	}
	return r
}

func main() {
	//STDOUT MUST BE FLUSHED MANUALLY
	defer writer.Flush()

	var n, x, y int
	scanf("%d\n", &n)
	scanf("%s\n", &s)
	scanf("%d %d\n", &x, &y)
	end = pos{x, y}
	printf("%d\n", solution())
}
