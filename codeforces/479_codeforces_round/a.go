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
	var n, k uint32
	fmt.Scan(&n, &k)
	for i := uint32(0); i < k; i++ {
		n = count(n)
	}
	fmt.Println(n)
}
