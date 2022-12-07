package fake

type FsElement interface {
	Size() int
	Name() string
	isDir() bool
}

type Directory struct {
	name        string
	content     []FsElement
	parent      *Directory
	cached_size int
}

type File struct {
	name string
	size int
}

type Path struct {
	current *Directory
	root    *Directory
}

func NewRoot() *Directory {
	r := new(Directory)
	r.name = "root"
	r.content = make([]FsElement, 0)
	r.parent = nil
	return r
}

func (d *Directory) Name() string {
	return d.name
}

func (d *Directory) Size() int {
	if d.cached_size == 0 {
		s := 0
		for _, e := range d.content {
			s += e.Size()
		}
		d.cached_size = s
	}
	return d.cached_size
}

func (d *Directory) Append(element FsElement) {
	d.content = append(d.content, element)
}

func (d *Directory) isDir() bool {
	return true
}

func (p *Path) Mkdir(name string) {
	d := new(Directory)
	d.name = name
	d.content = make([]FsElement, 0)
	d.parent = p.current
	p.current.Append(d)
}

func (p *Path) Touch(name string, size int) {
	f := new(File)
	f.name = name
	f.size = size
	p.current.Append(f)
}

func (p *Path) Cd(d *Directory) {
	p.current = d
}

func (f *File) Name() string {
	return f.name
}

func (f *File) Size() int {
	return f.size
}

func (f *File) isDir() bool {
	return false
}
