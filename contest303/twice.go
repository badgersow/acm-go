package main

import (
	"fmt"
)

func repeatedCharacter(s string) byte {
	var mask [256]int
	for _, c := range s {
		if mask[c] > 0 {
			return byte(c)
		}
		mask[c]++
	}
	return '?'
}

func main() {
	c := repeatedCharacter("abccbaacz")
	fmt.Println(string(c))
}
