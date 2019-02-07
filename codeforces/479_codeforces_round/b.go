package main

import (
	"fmt"
	// "reflect"
)

func count(num uint32) uint32 {
	if num%10 == 0 {
		num /= 10
	} else {
		num--
	}
	return num
}

func main() {
	var n byte
	max := 0
	counter := make(map[string]int)

	fmt.Scanf("%d\n", &n)
	input := make([]rune, n, 100)

	for i := byte(0); i < n; i++ {
		fmt.Scanf("%c", &input[i])
	}

	// fmt.Printf("%v", input)

	text := string(input)

	for i := 0; i < len(text)-1; i++ {
		pair := text[i : i+2]
		// fmt.Printf("%v\n", pair)
		counter[pair]++
		if counter[pair] > max {
			max = counter[pair]
		}
	}

	var answer string
	for k, v := range counter {
		if v == max {
			answer = k
			break
		}
	}

	fmt.Printf("%v\n", answer)
}
