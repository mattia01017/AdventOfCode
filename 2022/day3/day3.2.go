package main

import (
	"fmt"
	"io"
)

func priority(c byte) int {
	var priority int
	if c >= 'a' && c <= 'z' {
		priority = int(c - 'a' + 1)
	} else {
		priority = int(c - 'A' + 27)
	}
	return priority
}

func main() {
	var err error
	var input string
	var tot, nItems int

	for err != io.EOF {
		count := make(map[byte]int)
		for i := 0; i < 3; i++ {
			nItems, err = fmt.Scanln(&input)
			if nItems > 0 {
				seen := make(map[byte]bool)
				for j := 0; j < len(input); j++ {
					if !seen[input[j]] {
						seen[input[j]] = true
						count[input[j]]++
					}
					if count[input[j]] == 3 {
						tot += priority(input[j])
						break
					}
				}
			}
		}
	}
	fmt.Println(tot)
}
