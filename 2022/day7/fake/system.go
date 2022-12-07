package fake

import (
	"strconv"
	"strings"
)

type System struct {
	path *Path
}

func NewSystem() *System {
	i := new(System)
	i.path = new(Path)
	i.path.root = NewRoot()
	i.path.current = i.path.root
	return i
}

func (i *System) Read(command string) {
	if command[0] == '$' {
		i.exec(command[1:])
	} else {
		i.addToFS(command)
	}
}

func (i *System) addToFS(command string) {
	toks := strings.Fields(command)
	if toks[0] == "dir" {
		i.path.Mkdir(toks[1])
	} else {
		fsize, _ := strconv.Atoi(toks[0])
		fname := toks[1]
		i.path.Touch(fname, fsize)
	}
}

func (i *System) exec(command string) {
	toks := strings.Fields(command)
	if toks[0] == "cd" {
		switch toks[1] {
		case "..":
			i.path.Cd(i.path.current.parent)
		case "/":
			i.path.Cd(i.path.root)
		default:
			for _, e := range i.path.current.content {
				if e.isDir() && e.Name() == toks[1] {
					casted_e := e.(*Directory)
					i.path.Cd(casted_e)
				}
			}
		}
	}
}

func (i *System) SumOfSizes(sizeLimit int) int {
	return sumOfSizes(sizeLimit, i.path.root)
}

func sumOfSizes(sizeLimit int, current *Directory) int {
	var ret int
	for _, e := range current.content {
		if e.isDir() {
			if e.Size() <= sizeLimit {
				ret += e.Size()
			}
			casted_e := e.(*Directory)
			ret += sumOfSizes(sizeLimit, casted_e)
		}
	}
	return ret
}

func (i *System) UsedSpace() int {
	return i.path.root.Size()
}

func (i *System) FindSmallestDirGreaterThan(size int) *Directory {
	list := make([]*Directory, 0)
	findSmallestDirGreaterThan(size, i.path.root, &list)

	min := list[0]
	for _, dir := range list {
		if dir.Size() < min.Size() {
			min = dir
		}
	}
	return min
}

func findSmallestDirGreaterThan(size int, current *Directory, list *[]*Directory) {
	for _, e := range current.content {
		if e.isDir() {
			casted_e := e.(*Directory)
			if e.Size() >= size {
				*list = append(*list, casted_e)
			}
			findSmallestDirGreaterThan(size, casted_e, list)
		}
	}
}
