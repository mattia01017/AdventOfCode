package main

import (
	"fmt"
	"io"
)

func main() {
	var readVal,
		tempSum,
		numItems int
	var topSums [3]int
	var err error
	for err != io.EOF {
		numItems, err = fmt.Scanln(&readVal)
		if numItems == 0 {
			for i := 0; i < 3 && topSums[i] < tempSum; i++ {
				if i == 0 {
					topSums[i] = tempSum
				} else {
					topSums[i-1], topSums[i] = topSums[i], topSums[i-1]
				}
			}
			tempSum = 0
		} else {
			tempSum += readVal
		}
	}
	fmt.Printf("Migliori 3 elfi: %v, totale: ", topSums)
	tot := 0
	for _, val := range topSums {
		tot += val
	}
	fmt.Println(tot)
}
