package day07

type Filesystem struct {
	tld *Directory
	cur *Directory
	all []*Directory
}

type Directory struct {
	name   string
	parent *Directory
	dir    []*Directory
	flist  []File_t
	size   int
}

type File_t struct {
	name   string
	parent *Directory
	size   int
}

func newDirectory(name string, parent *Directory) *Directory {
	dir := Directory{
		name:   name,
		parent: parent,
		size:   0,
	}

	return &dir
}

func (dir *Directory) addDirectory(newDir Directory) {
	dir.dir = append(dir.dir, &newDir)
	// lib.Print("%p: %v", &newDir, newDir.getPath())
}

func (fs *Filesystem) addDirectory(newDir Directory) {
	fs.all = append(fs.all, &newDir)
}

func (dir *Directory) addFile(newFile File_t) {
	dir.flist = append(dir.flist, newFile)

	scandir := dir
	for scandir != nil {
		scandir.size += newFile.size
		scandir = scandir.parent
	}
}

func (dir *Directory) getPath() string {
	var path []string

	scandir := dir
	for scandir != nil {
		path = append(path, scandir.name)
		scandir = scandir.parent
	}

	strpath := ""
	for i := len(path) - 1; i >= 0; i -= 1 {
		strpath += path[i]
		if i > 0 && i < len(path)-1 {
			strpath += "/"
		}
	}
	return strpath
}

func (dir *Directory) getDirectory(name string) *Directory {
	for _, d := range dir.dir {
		if d.name == name {
			return d
		}
	}
	return nil
}

func newFile(name string, size int, parent *Directory) *File_t {
	file := File_t{
		name:   name,
		parent: parent,
		size:   size,
	}

	return &file
}
