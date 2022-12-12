package main

import (
	"day12/grid"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	input_b, _ := ioutil.ReadAll(os.Stdin)
	input := string(input_b)

	g := grid.NewGrid(input[:len(input)-1])

	fmt.Println("Cammino minimo da S:", g.ShortestPathLen())
	fmt.Println("Cammino minimo da un qualsiasi punto 'a':", g.ShortestPathLenNoStart())
}
