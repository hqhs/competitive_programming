package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func gcd(a, b int) int {
	var t int
	for b != 0 {
		t = b
		b = a % b
		a = t
	}
	return a
}

func main() {
	//STDOUT MUST BE FLUSHED MANUALLY
	defer writer.Flush()
	strs := []string{
		"0000", "0001", "0010", "0011",
		"0100", "0101", "0110", "0111",
		"1000", "1001", "1010", "1011",
		"1100", "1101", "1110", "1111",
	}
	bin := make([][]byte, 16)
	for i, val := range strs {
		bin[i] = []byte(val)
	}

	r := bufio.NewReader(reader)
	// first line: read n
	var n int
	l, _ := r.ReadString(byte('\n'))
	fmt.Sscanf(l, "%d\n", &n)
	var t []byte
	m := make([][]byte, n)
	for i := 0; i < n; i++ {
		tmp := make([]byte, 0, n)
		// t is just a pointer, we need to convert/copy it
		// skip whitespace
		for {
			if c, _, _ := r.ReadRune(); !unicode.IsSpace(c) {
				r.UnreadRune()
				break
			}
		}
		t, _ = r.ReadSlice(byte('\n'))
		// remove newline
		t = t[:len(t)-1]
		// 0 = 48, 9 = 57, A = 65, F = 70
		for _, c := range t {
			if 48 <= c && c <= 57 {
				tmp = append(tmp, bin[c-48]...)
			} else if 65 <= c && c <= 70 {
				tmp = append(tmp, bin[c-55]...)
			}
		}
		m[i] = tmp
	}

	// solution
	g := n
	// iterate over rows
	for i := range m {
		for j := 0; j < n; j++ {
			k := j
			for k < n && m[i][j] == m[i][k] {
				k++
			}
			g = gcd(g, k-j)
			j = k - 1
		}
	}

	// iterate over columns
	for j := 0; j < n; j++ {
		for i := 0; i < n; i++ {
			k := i
			for k < n && m[k][j] == m[i][j] {
				k++
			}
			g = gcd(g, k-i)
			i = k - 1
		}
	}

	printf("%v\n", g)
}
