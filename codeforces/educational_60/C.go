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

type p struct {
	x int64
	y int64
}

func abs(x int64) int64 {
	if x > 0 {
		return x
	}
	return -x
}

func (p *p) String() string {
	return fmt.Sprintf("x: %v; y: %v\n", p.x, p.y)
}

func main() {
	//STDOUT MUST BE FLUSHED MANUALLY
	defer writer.Flush()
	var n int64
	start := p{0, 0}
	end := p{0, 0}
	scanf("%d %d\n", &start.x, &start.y)
	scanf("%d %d\n", &end.x, &end.y)
	scanf("%d\n", &n)
	var forecast string
	scanf("%s\n", &forecast)
	move := make([]p, n+1)
	cur := p{0, 0} // cur is endpoint where ship is going to be if it stays still
	move[0] = cur
	for i, c := range forecast {
		switch c {
		case 'U':
			cur.y++
		case 'D':
			cur.y--
		case 'L':
			cur.x--
		case 'R':
			cur.x++
		}
		move[i+1] = cur
	}
	// printf("move: %+v\n", move)
	// printf("end: %+v\n", end)
	// Let's use binary search. K is amount of days used for journey
	// To get to finish point we need to move for atleas 1 time close
	// to endpoint. Max manchaten distance is 2 * 10*9, max weather forecast
	// length is 10^5, hence upper bound is 2 * 10 ^ 14.
	// lower bound is 0
	// To check what (x2, y2) is accesible, lets calculate (x3, y3)
	// k value for (x3,y3) found with binary search, and ship stays still
	// if |x2 - x3| + |y2 - y3| <= k then (x2,y2) is accesible
	// To calculate (x3,y3) we know what k/n cycles past and k % n days of new cicle
	// Whats O(1) using move array above

	var x, y, dist int64
	l, r := int64(0), int64(1e18)
	for (r - l) > 1 {
		mid := (int64(l+r) / 2)
		cnt, rem := mid/n, mid%n
		x = start.x + move[rem].x + cnt*move[n].x
		y = start.y + move[rem].y + cnt*move[n].y
		dist = abs(x-end.x) + abs(y-end.y)
		if dist <= mid {
			r = mid
		} else {
			l = mid
		}
	}

	if r > 5e17 {
		r = int64(-1)
	}

	printf("%d\n", r)
}
