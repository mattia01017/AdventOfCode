package main

import (
	"fmt"
	"io"
)

func main() {
	var err error
	var input string
	var totPriority,
		nItems int
	for err != io.EOF {
		nItems, err = fmt.Scanln(&input)
		first := input[:len(input)/2]
		second := input[len(input)/2:]
		if nItems > 0 {
			var charFirst,
				charSecond byte
			for i := 0; i < len(first); i++ {
				charFirst = first[i]
				j := 0
				for j = 0; j < len(second); j++ {
					charSecond = second[j]
					if charFirst == charSecond {
						if charFirst >= 'a' && charFirst <= 'z' {
							totPriority += int(charFirst - 'a' + 1)
						} else {
							totPriority += int(charFirst - 'A' + 27)
						}
						break
					}
				}
				if j < len(second) {
					break
				}
			}
		}
	}
	fmt.Printf("Totale: %d\n", totPriority)
}
