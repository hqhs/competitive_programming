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

func checkDiverce(s string) string {
	if len(s) < 2 {
		return ""
	}
	for index, char := range s[1:] {
		// if previous char != current, return previous+current
		if s[index] != byte(char) {
			return fmt.Sprintf("%c%c", s[index], char)
		}
	}
	return ""
}

func main() {
	//STDOUT MUST BE FLUSHED MANUALLY
	defer writer.Flush()

	// var a, b int
	var input string
	var n int
	scanf("%d\n", &n)
	scanf("%s", &input)
	if answer := checkDiverce(input); answer != "" {
		printf("YES\n%s\n", answer)
	} else {
		printf("NO\n")
	}
}
