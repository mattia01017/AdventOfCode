package main

import (
	"bufio"
	"day7/fake"
	"fmt"
	"io"
	"os"
)

const totalSpace = 70000000
const neededSpace = 30000000

func main() {
	i := fake.NewSystem()
	r := bufio.NewReader(os.Stdin)
	var err error
	var line string
	for true {
		line, err = r.ReadString('\n')
		if err == io.EOF {
			break
		}
		i.Read(line)
	}
	fmt.Println("Sum of sizes of files smaller than 100000:", i.SumOfSizes(100000))
	usedSpace := i.UsedSpace()
	freeSpace := totalSpace - usedSpace
	fmt.Println("Total space:", totalSpace)
	fmt.Println("Used space:", usedSpace)
	fmt.Println("Free space:", freeSpace)
	toFree := neededSpace - freeSpace
	toErase := i.SmallestDirToErase(toFree)
	fmt.Printf("Smallest directory to erase to free enought space: %s, size %d\n", toErase.Name(), toErase.Size())
}
